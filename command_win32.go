//go:build windows
// Code generated by go-vk from vk-1.2.203.xml at 2023-03-28 13:29:28.4015828 -0500 CDT m=+3.316157401. DO NOT EDIT.

package vk

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// AcquireFullScreenExclusiveModeEXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkAcquireFullScreenExclusiveModeEXT.html
func AcquireFullScreenExclusiveModeEXT(device Device, swapchain SwapchainKHR) (r error) {

	r = Result(execTrampoline(vkAcquireFullScreenExclusiveModeEXT, uintptr(device), uintptr(swapchain)))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkAcquireFullScreenExclusiveModeEXT = &vkCommand{"vkAcquireFullScreenExclusiveModeEXT", 2, true, nil}

// AcquireWinrtDisplayNV: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkAcquireWinrtDisplayNV.html
func AcquireWinrtDisplayNV(physicalDevice PhysicalDevice, display DisplayKHR) (r error) {

	r = Result(execTrampoline(vkAcquireWinrtDisplayNV, uintptr(physicalDevice), uintptr(display)))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkAcquireWinrtDisplayNV = &vkCommand{"vkAcquireWinrtDisplayNV", 2, true, nil}

// CreateWin32SurfaceKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateWin32SurfaceKHR.html
func CreateWin32SurfaceKHR(instance Instance, createInfo *Win32SurfaceCreateInfoKHR, allocator *AllocationCallbacks) (surface SurfaceKHR, r error) {
	// Parameter is a singular input, requires translation - createInfo
	var pCreateInfo *_vkWin32SurfaceCreateInfoKHR
	if createInfo != nil {
		pCreateInfo = createInfo.Vulkanize()
	}

	// Parameter is a singular input, pass direct - allocator
	var pAllocator unsafe.Pointer
	if allocator != nil {
		pAllocator = unsafe.Pointer(allocator)
	}

	// surface is a binding-allocated single return value and will be populated by Vulkan
	ptr_pSurface := &surface

	r = Result(execTrampoline(vkCreateWin32SurfaceKHR, uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(ptr_pSurface))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkCreateWin32SurfaceKHR = &vkCommand{"vkCreateWin32SurfaceKHR", 4, true, nil}

// GetDeviceGroupSurfacePresentModes2EXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModes2EXT.html
func GetDeviceGroupSurfacePresentModes2EXT(device Device, surfaceInfo *PhysicalDeviceSurfaceInfo2KHR) (modes DeviceGroupPresentModeFlagsKHR, r error) {
	// Parameter is a singular input, requires translation - surfaceInfo
	var pSurfaceInfo *_vkPhysicalDeviceSurfaceInfo2KHR
	if surfaceInfo != nil {
		pSurfaceInfo = surfaceInfo.Vulkanize()
	}

	// modes is a binding-allocated single return value and will be populated by Vulkan
	ptr_pModes := &modes

	r = Result(execTrampoline(vkGetDeviceGroupSurfacePresentModes2EXT, uintptr(device), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(ptr_pModes))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkGetDeviceGroupSurfacePresentModes2EXT = &vkCommand{"vkGetDeviceGroupSurfacePresentModes2EXT", 3, true, nil}

// GetFenceWin32HandleKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetFenceWin32HandleKHR.html
func GetFenceWin32HandleKHR(device Device, getWin32HandleInfo *FenceGetWin32HandleInfoKHR) (handle windows.Handle, r error) {
	// Parameter is a singular input, requires translation - getWin32HandleInfo
	var pGetWin32HandleInfo *_vkFenceGetWin32HandleInfoKHR
	if getWin32HandleInfo != nil {
		pGetWin32HandleInfo = getWin32HandleInfo.Vulkanize()
	}

	// handle is a binding-allocated single return value and will be populated by Vulkan
	ptr_pHandle := &handle

	r = Result(execTrampoline(vkGetFenceWin32HandleKHR, uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(ptr_pHandle))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkGetFenceWin32HandleKHR = &vkCommand{"vkGetFenceWin32HandleKHR", 3, true, nil}

// GetMemoryWin32HandleKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandleKHR.html
func GetMemoryWin32HandleKHR(device Device, getWin32HandleInfo *MemoryGetWin32HandleInfoKHR) (handle windows.Handle, r error) {
	// Parameter is a singular input, requires translation - getWin32HandleInfo
	var pGetWin32HandleInfo *_vkMemoryGetWin32HandleInfoKHR
	if getWin32HandleInfo != nil {
		pGetWin32HandleInfo = getWin32HandleInfo.Vulkanize()
	}

	// handle is a binding-allocated single return value and will be populated by Vulkan
	ptr_pHandle := &handle

	r = Result(execTrampoline(vkGetMemoryWin32HandleKHR, uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(ptr_pHandle))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkGetMemoryWin32HandleKHR = &vkCommand{"vkGetMemoryWin32HandleKHR", 3, true, nil}

// GetMemoryWin32HandleNV: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandleNV.html
func GetMemoryWin32HandleNV(device Device, memory DeviceMemory, handleType ExternalMemoryHandleTypeFlagsNV) (handle windows.Handle, r error) {
	// handle is a binding-allocated single return value and will be populated by Vulkan
	ptr_pHandle := &handle

	r = Result(execTrampoline(vkGetMemoryWin32HandleNV, uintptr(device), uintptr(memory), uintptr(handleType), uintptr(unsafe.Pointer(ptr_pHandle))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkGetMemoryWin32HandleNV = &vkCommand{"vkGetMemoryWin32HandleNV", 4, true, nil}

// GetMemoryWin32HandlePropertiesKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandlePropertiesKHR.html
func GetMemoryWin32HandlePropertiesKHR(device Device, handleType ExternalMemoryHandleTypeFlagBits, handle windows.Handle) (memoryWin32HandleProperties MemoryWin32HandlePropertiesKHR, r error) {
	// memoryWin32HandleProperties is a binding-allocated single return value and will be populated by Vulkan, but requiring translation
	var pMemoryWin32HandleProperties *_vkMemoryWin32HandlePropertiesKHR = memoryWin32HandleProperties.Vulkanize()

	r = Result(execTrampoline(vkGetMemoryWin32HandlePropertiesKHR, uintptr(device), uintptr(handleType), uintptr(handle), uintptr(unsafe.Pointer(pMemoryWin32HandleProperties))))

	memoryWin32HandleProperties = *(pMemoryWin32HandleProperties.Goify())
	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkGetMemoryWin32HandlePropertiesKHR = &vkCommand{"vkGetMemoryWin32HandlePropertiesKHR", 4, true, nil}

// GetPhysicalDeviceSurfacePresentModes2EXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html
func GetPhysicalDeviceSurfacePresentModes2EXT(physicalDevice PhysicalDevice, surfaceInfo *PhysicalDeviceSurfaceInfo2KHR) (presentModes []PresentModeKHR, r error) {
	// Parameter is a singular input, requires translation - surfaceInfo
	var pSurfaceInfo *_vkPhysicalDeviceSurfaceInfo2KHR
	if surfaceInfo != nil {
		pSurfaceInfo = surfaceInfo.Vulkanize()
	}

	// presentModes is a double-call array output
	var presentModeCount uint32
	pPresentModeCount := &presentModeCount
	// first trampoline happens here; also, still need to check returned Result value
	// NOT identical internal and external, result needs translation
	var pPresentModes *PresentModeKHR

	r = Result(execTrampoline(vkGetPhysicalDeviceSurfacePresentModes2EXT, uintptr(physicalDevice), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pPresentModeCount)), uintptr(unsafe.Pointer(pPresentModes))))

	sl_pPresentModes := make([]PresentModeKHR, presentModeCount)
	presentModes = make([]PresentModeKHR, presentModeCount)
	pPresentModes = &sl_pPresentModes[0]

	// Trampoline call after last array allocation
	r = Result(execTrampoline(vkGetPhysicalDeviceSurfacePresentModes2EXT, uintptr(physicalDevice), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pPresentModeCount)), uintptr(unsafe.Pointer(pPresentModes))))

	for i := range sl_pPresentModes {
		presentModes[i] = *&sl_pPresentModes[i]
	}
	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkGetPhysicalDeviceSurfacePresentModes2EXT = &vkCommand{"vkGetPhysicalDeviceSurfacePresentModes2EXT", 4, true, nil}

// GetPhysicalDeviceWin32PresentationSupportKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceWin32PresentationSupportKHR.html
func GetPhysicalDeviceWin32PresentationSupportKHR(physicalDevice PhysicalDevice, queueFamilyIndex uint32) (r bool) {

	rval := Bool32(execTrampoline(vkGetPhysicalDeviceWin32PresentationSupportKHR, uintptr(physicalDevice), uintptr(queueFamilyIndex)))
	r = translatePublic_Bool32(rval)

	return
}

var vkGetPhysicalDeviceWin32PresentationSupportKHR = &vkCommand{"vkGetPhysicalDeviceWin32PresentationSupportKHR", 2, true, nil}

// GetSemaphoreWin32HandleKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreWin32HandleKHR.html
func GetSemaphoreWin32HandleKHR(device Device, getWin32HandleInfo *SemaphoreGetWin32HandleInfoKHR) (handle windows.Handle, r error) {
	// Parameter is a singular input, requires translation - getWin32HandleInfo
	var pGetWin32HandleInfo *_vkSemaphoreGetWin32HandleInfoKHR
	if getWin32HandleInfo != nil {
		pGetWin32HandleInfo = getWin32HandleInfo.Vulkanize()
	}

	// handle is a binding-allocated single return value and will be populated by Vulkan
	ptr_pHandle := &handle

	r = Result(execTrampoline(vkGetSemaphoreWin32HandleKHR, uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(ptr_pHandle))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkGetSemaphoreWin32HandleKHR = &vkCommand{"vkGetSemaphoreWin32HandleKHR", 3, true, nil}

// GetWinrtDisplayNV: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetWinrtDisplayNV.html
func GetWinrtDisplayNV(physicalDevice PhysicalDevice, deviceRelativeId uint32) (display DisplayKHR, r error) {
	// display is a binding-allocated single return value and will be populated by Vulkan
	ptr_pDisplay := &display

	r = Result(execTrampoline(vkGetWinrtDisplayNV, uintptr(physicalDevice), uintptr(deviceRelativeId), uintptr(unsafe.Pointer(ptr_pDisplay))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkGetWinrtDisplayNV = &vkCommand{"vkGetWinrtDisplayNV", 3, true, nil}

// ImportFenceWin32HandleKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkImportFenceWin32HandleKHR.html
func ImportFenceWin32HandleKHR(device Device, importFenceWin32HandleInfo *ImportFenceWin32HandleInfoKHR) (r error) {
	// Parameter is a singular input, requires translation - importFenceWin32HandleInfo
	var pImportFenceWin32HandleInfo *_vkImportFenceWin32HandleInfoKHR
	if importFenceWin32HandleInfo != nil {
		pImportFenceWin32HandleInfo = importFenceWin32HandleInfo.Vulkanize()
	}

	r = Result(execTrampoline(vkImportFenceWin32HandleKHR, uintptr(device), uintptr(unsafe.Pointer(pImportFenceWin32HandleInfo))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkImportFenceWin32HandleKHR = &vkCommand{"vkImportFenceWin32HandleKHR", 2, true, nil}

// ImportSemaphoreWin32HandleKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkImportSemaphoreWin32HandleKHR.html
func ImportSemaphoreWin32HandleKHR(device Device, importSemaphoreWin32HandleInfo *ImportSemaphoreWin32HandleInfoKHR) (r error) {
	// Parameter is a singular input, requires translation - importSemaphoreWin32HandleInfo
	var pImportSemaphoreWin32HandleInfo *_vkImportSemaphoreWin32HandleInfoKHR
	if importSemaphoreWin32HandleInfo != nil {
		pImportSemaphoreWin32HandleInfo = importSemaphoreWin32HandleInfo.Vulkanize()
	}

	r = Result(execTrampoline(vkImportSemaphoreWin32HandleKHR, uintptr(device), uintptr(unsafe.Pointer(pImportSemaphoreWin32HandleInfo))))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkImportSemaphoreWin32HandleKHR = &vkCommand{"vkImportSemaphoreWin32HandleKHR", 2, true, nil}

// ReleaseFullScreenExclusiveModeEXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkReleaseFullScreenExclusiveModeEXT.html
func ReleaseFullScreenExclusiveModeEXT(device Device, swapchain SwapchainKHR) (r error) {

	r = Result(execTrampoline(vkReleaseFullScreenExclusiveModeEXT, uintptr(device), uintptr(swapchain)))

	if r == Result(0) {
		r = SUCCESS
	}
	return
}

var vkReleaseFullScreenExclusiveModeEXT = &vkCommand{"vkReleaseFullScreenExclusiveModeEXT", 2, true, nil}
