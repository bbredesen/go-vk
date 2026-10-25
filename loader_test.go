package vk

import (
	"os"
	"testing"
	"unsafe"
)

// Integration tests against a real Vulkan loader. They are opt-in: set VK_TEST_LOADER to the path of the loader library
// (for example the Vulkan SDK's libvulkan.dylib, libvulkan.so, or vulkan-1.dll) and select them with -run Loader:
//
//	VK_TEST_LOADER=/usr/local/lib/libvulkan.dylib go test -run Loader -v
//
// When VK_TEST_LOADER is set, the fake-library tests are skipped, because the process can only load one Vulkan library.
//
// Background: the Khronos loader exports only core commands plus the surface, swapchain and display extensions.
// Every other extension command (VK_EXT_debug_utils, VK_KHR_push_descriptor, ...) must be obtained from
// vkGetInstanceProcAddr or vkGetDeviceProcAddr, and is not found by dlsym/GetProcAddress; see resolveSlow() in
// static_common.go. A physical device is not required: these tests need only a loader and an instance.

func requireLoader(t *testing.T) {
	t.Helper()
	if os.Getenv(loaderEnv) == "" {
		t.Skip("set " + loaderEnv + " to a Vulkan loader library to run the loader integration tests")
	}
}

// tryResolve resolves a command's function pointer. It reports failure either as a nil pointer or as a panic, so the test
// remains valid however resolution reports a missing command.
func tryResolve(cmd *vkCommand) (fn unsafe.Pointer, panicked any) {
	defer func() { panicked = recover() }()
	return cmd.resolve(), nil
}

// newTestInstance creates an instance with VK_EXT_debug_utils (and portability enumeration, where offered, so that
// MoltenVK is visible to the loader). It skips the test if the loader does not offer debug utils.
func newTestInstance(t *testing.T) Instance {
	t.Helper()

	available, r := EnumerateInstanceExtensionProperties("")
	if r != nil && r != SUCCESS {
		t.Fatalf("EnumerateInstanceExtensionProperties: %v", r)
	}
	offers := func(name string) bool {
		for _, p := range available {
			if p.ExtensionName == name {
				return true
			}
		}
		return false
	}

	if !offers(EXT_DEBUG_UTILS_EXTENSION_NAME) {
		t.Skip(EXT_DEBUG_UTILS_EXTENSION_NAME + " is not offered by this loader")
	}

	info := &InstanceCreateInfo{PpEnabledExtensionNames: []string{EXT_DEBUG_UTILS_EXTENSION_NAME}}
	if offers("VK_KHR_portability_enumeration") {
		const enumeratePortabilityBit = 0x1 // VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR
		info.Flags = InstanceCreateFlags(enumeratePortabilityBit)
		info.PpEnabledExtensionNames = append(info.PpEnabledExtensionNames, "VK_KHR_portability_enumeration")
	}

	instance, r := CreateInstance(info, nil)
	if r != nil && r != SUCCESS {
		t.Fatalf("CreateInstance: %v", r)
	}
	t.Cleanup(func() { DestroyInstance(instance, nil) })
	return instance
}

func TestLoaderResolvesCommands(t *testing.T) {
	requireLoader(t)
	newTestInstance(t) // commands may only be resolvable once an instance exists

	tests := []struct {
		name string
		cmd  *vkCommand
	}{
		// Controls: exported by the loader, so found by dlsym today
		{"core instance command", vkEnumeratePhysicalDevices},
		{"core device command", vkCmdSetLineWidth},

		// Extension commands the loader does not export; they require vkGetInstanceProcAddr
		{"VK_EXT_debug_utils create", vkCreateDebugUtilsMessengerEXT},
		{"VK_EXT_debug_utils destroy", vkDestroyDebugUtilsMessengerEXT},
		{"VK_EXT_debug_utils submit", vkSubmitDebugUtilsMessageEXT},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fn, panicked := tryResolve(tc.cmd)
			if panicked != nil {
				t.Fatalf("resolving %s panicked: %v", tc.cmd.protoName, panicked)
			}
			if fn == nil {
				t.Fatalf("%s could not be resolved", tc.cmd.protoName)
			}
		})
	}
}

// TestLoaderCallsExtensionCommand calls an extension command end to end. With no messengers registered the loader
// accepts the message and does nothing, so the call must simply return.
func TestLoaderCallsExtensionCommand(t *testing.T) {
	requireLoader(t)
	instance := newTestInstance(t)

	// Never call through an unresolved pointer: that is a fatal signal, not a test failure
	if fn, panicked := tryResolve(vkSubmitDebugUtilsMessageEXT); panicked != nil {
		t.Fatalf("resolving %s panicked: %v", vkSubmitDebugUtilsMessageEXT.protoName, panicked)
	} else if fn == nil {
		t.Fatalf("%s could not be resolved", vkSubmitDebugUtilsMessageEXT.protoName)
	}

	SubmitDebugUtilsMessageEXT(instance,
		DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT,
		DebugUtilsMessageTypeFlagsEXT(DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT),
		&DebugUtilsMessengerCallbackDataEXT{PMessageIdName: "vk-gen", PMessage: "loader integration test"})
}
