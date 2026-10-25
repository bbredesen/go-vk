package vk

import (
	"bytes"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"
)

// #include <stdlib.h>
// #include "dlload.h"
import "C"

// Vulkanizer allows conversion from go-vk style structs to Vulkan-native structs. This
// includes setting the structure type flag, converting slices to pointers, etc.
type Vulkanizer interface {
	Vulkanize() unsafe.Pointer
}

// Goifier converts Vulkan-native structs back into go-vk style structs
type Goifier interface {
	Goify() Vulkanizer
}

// max is an internal utility function, used in processing struct member slice/array lengths
func max(nums ...int) int {
	rval := 0
	for _, v := range nums {
		if rval < v {
			rval = v
		}
	}
	return rval
}

// Error implements the error interface
// TODO: A way for commands to indicate if the Result code is an error for that command, or an unexpected return value?
func (r Result) Error() string {
	return r.String()
}

type vkCommand struct {
	protoName string
	argCount  int
	hasReturn bool
	fnHandle  unsafe.Pointer // read and written atomically; see resolve()
	aliases   []string       // other names Vulkan knows this command by; see resolveSlow()
}

var dlHandle unsafe.Pointer

var overrideLibName string

// OverrideDefaultVulkanLibrary allows you to set a specific Vulkan library name to be used in your program. For
// example, if you want to enable the validation layers, those layers are only available in the Vulkan SDK libary. go-vk
// passes the name to the host operating system's library opening/search method, so you must provide a relative or
// absolute path if your Vulkan library is not in the default search path for the platform.
func OverrideDefaultVulkanLibrary(nameOrPath string) {
	overrideLibName = nameOrPath
}

// Command resolution
//
// A command's function pointer is found the first time it is called, in this order:
//
//  1. vkGetInstanceProcAddr(NULL, name), which returns the global commands (vkCreateInstance and friends)
//  2. vkGetInstanceProcAddr(instance, name), using a live instance created through CreateInstance. This is the only way
//     to find extension commands through the Khronos loader, which does not export them as symbols. Through the loader,
//     device-level commands resolved this way are trampolines that dispatch on the handle passed to them, so they work
//     for any device of the instance, and no per-device state is needed.
//  3. dlsym/GetProcAddress on the library, for libraries that export commands directly (such as MoltenVK) or that do not
//     export vkGetInstanceProcAddr
//
// If the command's own name is found by none of those, the names it is aliased by (vkCommand.aliases)
// are tried in the same way, because a device that has an extension but not the core version that absorbed it offers only
// the extension's name for the command.
//
// A command that is found nowhere panics with its name. Callers are expected to have checked, through the Vulkan API,
// that the extension providing a command is available and enabled before calling it.
//
// Resolved pointers are cleared when an instance is destroyed (DestroyInstance), because the specification ties them to
// the instance that produced them. resolveMu guards the state below, and the fast path in resolve() takes no lock.
var (
	resolveMu     sync.Mutex
	getProcAddr   unsafe.Pointer // vkGetInstanceProcAddr, or nil if the library does not export it
	liveInstances []Instance     // in creation order; the most recent is used to resolve
	resolvedCmds  []*vkCommand   // commands with a cached pointer, so that it can be cleared
)

// instanceCreated and instanceDestroyed are called from the generated CreateInstance and DestroyInstance (see the "command"
// section of exceptions.json).
func instanceCreated(instance Instance) {
	resolveMu.Lock()
	defer resolveMu.Unlock()

	liveInstances = append(liveInstances, instance)
}

func instanceDestroyed(instance Instance) {
	resolveMu.Lock()
	defer resolveMu.Unlock()

	for i, live := range liveInstances {
		if live == instance {
			liveInstances = append(liveInstances[:i], liveInstances[i+1:]...)
			break
		}
	}

	for _, cmd := range resolvedCmds {
		atomic.StorePointer(&cmd.fnHandle, nil)
	}
	resolvedCmds = nil
}

// resolve returns the command's function pointer, finding it on first use.
func (cmd *vkCommand) resolve() unsafe.Pointer {
	if fn := atomic.LoadPointer(&cmd.fnHandle); fn != nil {
		return fn
	}
	return cmd.resolveSlow()
}

func (cmd *vkCommand) resolveSlow() unsafe.Pointer {
	resolveMu.Lock()
	defer resolveMu.Unlock()

	if fn := atomic.LoadPointer(&cmd.fnHandle); fn != nil {
		return fn // another goroutine resolved it while this one waited for the lock
	}

	openLibrary()

	var fn unsafe.Pointer
	if fn = lookupCommand(cmd.protoName); fn == nil {
		// A command promoted to core, or to another extension, keeps its old name as an alias, and a device that
		// supports the extension but not the newer version offers only that name. The signatures are identical.
		for _, alias := range cmd.aliases {
			if fn = lookupCommand(alias); fn != nil {
				break
			}
		}
	}
	if fn == nil {
		msg := "vk: " + cmd.protoName + " is not available: the extension that provides it is not supported or not enabled " +
			"(or, for an instance or device level command, no instance has been created)"
		if len(cmd.aliases) > 0 {
			msg += "; also tried " + strings.Join(cmd.aliases, ", ")
		}
		panic(msg)
	}

	atomic.StorePointer(&cmd.fnHandle, fn)
	resolvedCmds = append(resolvedCmds, cmd)
	return fn
}

// lookupCommand finds one command name: through vkGetInstanceProcAddr with no instance, then with the newest live
// instance, then the library's own exports (the order described under "Command resolution" above). It returns nil if the
// name is not found. The caller holds resolveMu and has opened the library.
func lookupCommand(name string) unsafe.Pointer {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var fn unsafe.Pointer
	if getProcAddr != nil {
		fn = C.CallGetProcAddr(getProcAddr, 0, cname)
		if fn == nil && len(liveInstances) > 0 {
			fn = C.CallGetProcAddr(getProcAddr, C.uintptr_t(liveInstances[len(liveInstances)-1]), cname)
		}
	}
	if fn == nil {
		fn = C.SymbolFromName(dlHandle, unsafe.Pointer(cname))
	}
	return fn
}

// openLibrary loads the Vulkan library on first use. The caller must hold resolveMu.
func openLibrary() {
	if dlHandle != nil {
		return
	}

	var libName string
	switch runtime.GOOS {
	case "windows":
		libName = "vulkan-1.dll"
	case "darwin":
		libName = "libMoltenVK.dylib"
	case "linux":
		// TODO: Running on Linux is tested only to the point of creating and
		// destroying a Vulkan instance.
		libName = "libvulkan.so"
	default:
		panic("Unsupported GOOS at OpenLibrary: " + runtime.GOOS)
	}

	if overrideLibName != "" {
		libName = overrideLibName
	}

	cstr := C.CString(libName)
	defer C.free(unsafe.Pointer(cstr))

	dlHandle = C.OpenLibrary(cstr)
	if dlHandle == nil {
		panic("vk: could not load the Vulkan library " + libName)
	}

	name := C.CString("vkGetInstanceProcAddr")
	defer C.free(unsafe.Pointer(name))
	getProcAddr = C.SymbolFromName(dlHandle, unsafe.Pointer(name))
}

func stringToNullTermBytes(s string) *byte {
	b := []byte(s)
	b = append(b, 0)
	return &b[0]
}

func nullTermBytesToString(b []byte) string {
	n := bytes.IndexByte(b, 0)
	return string(b[:n])
}

// cStringToString copies a null-terminated C string into a Go string. A nil pointer returns the empty string.
func cStringToString(p unsafe.Pointer) string {
	if p == nil {
		return ""
	}
	n := 0
	for *(*byte)(unsafe.Add(p, n)) != 0 {
		n++
	}
	return string(unsafe.Slice((*byte)(p), n))
}
