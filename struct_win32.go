//go:build windows
// Code generated by go-vk from vk.xml at 2023-03-07 23:39:05.95336 +0100 CET m=+2.845645690. DO NOT EDIT.

package vk

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// D3D12FenceSubmitInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkD3D12FenceSubmitInfoKHR.html
type D3D12FenceSubmitInfoKHR struct {
	// SType = STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR
	PNext unsafe.Pointer
	// waitSemaphoreValuesCount
	PWaitSemaphoreValues []uint64
	// signalSemaphoreValuesCount
	PSignalSemaphoreValues []uint64
}

type _vkD3D12FenceSubmitInfoKHR struct {
	sType                      StructureType
	pNext                      unsafe.Pointer
	waitSemaphoreValuesCount   uint32
	pWaitSemaphoreValues       *uint64
	signalSemaphoreValuesCount uint32
	pSignalSemaphoreValues     *uint64
}

func (s *_vkD3D12FenceSubmitInfoKHR) Goify() *D3D12FenceSubmitInfoKHR {
	rval := &D3D12FenceSubmitInfoKHR{
		PNext: (unsafe.Pointer)(s.pNext),
		// Unexpected 'isLenForAnotherMember'!
		// Unexpected 'isLenForAnotherMember'!
	}
	return rval
}
func (s *D3D12FenceSubmitInfoKHR) Vulkanize() *_vkD3D12FenceSubmitInfoKHR {
	if s == nil {
		return nil
	}

	var psl_pWaitSemaphoreValues *uint64
	if len(s.PWaitSemaphoreValues) > 0 {
		psl_pWaitSemaphoreValues = &s.PWaitSemaphoreValues[0]
	}

	var psl_pSignalSemaphoreValues *uint64
	if len(s.PSignalSemaphoreValues) > 0 {
		psl_pSignalSemaphoreValues = &s.PSignalSemaphoreValues[0]
	}
	rval := &_vkD3D12FenceSubmitInfoKHR{
		sType:                      STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR, /*c1*/
		pNext:                      (unsafe.Pointer)(s.PNext),                  /*cb*/
		waitSemaphoreValuesCount:   uint32(len(s.PWaitSemaphoreValues)),        /*c6-a*/
		pWaitSemaphoreValues:       psl_pWaitSemaphoreValues,                   /*c rem*/
		signalSemaphoreValuesCount: uint32(len(s.PSignalSemaphoreValues)),      /*c6-a*/
		pSignalSemaphoreValues:     psl_pSignalSemaphoreValues,                 /*c rem*/
	}
	return rval
}

// ExportFenceWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkExportFenceWin32HandleInfoKHR.html
type ExportFenceWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR
	PNext       unsafe.Pointer
	PAttributes *windows.SecurityAttributes
	DwAccess    uint32
	Name        unsafe.Pointer
}

type _vkExportFenceWin32HandleInfoKHR struct {
	sType       StructureType
	pNext       unsafe.Pointer
	pAttributes *windows.SecurityAttributes
	dwAccess    uint32
	name        unsafe.Pointer
}

func (s *_vkExportFenceWin32HandleInfoKHR) Goify() *ExportFenceWin32HandleInfoKHR {
	rval := &ExportFenceWin32HandleInfoKHR{
		PNext:       (unsafe.Pointer)(s.pNext),
		PAttributes: (*windows.SecurityAttributes)(s.pAttributes),
		DwAccess:    (uint32)(s.dwAccess),
		Name:        (unsafe.Pointer)(s.name),
	}
	return rval
}
func (s *ExportFenceWin32HandleInfoKHR) Vulkanize() *_vkExportFenceWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkExportFenceWin32HandleInfoKHR{
		sType:       STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR, /*c1*/
		pNext:       (unsafe.Pointer)(s.PNext),                         /*cb*/
		pAttributes: (*windows.SecurityAttributes)(s.PAttributes),      /*cb*/
		dwAccess:    (uint32)(s.DwAccess),                              /*cb*/
		name:        (unsafe.Pointer)(s.Name),                          /*cb*/
	}
	return rval
}

// ExportMemoryWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkExportMemoryWin32HandleInfoKHR.html
type ExportMemoryWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR
	PNext       unsafe.Pointer
	PAttributes *windows.SecurityAttributes
	DwAccess    uint32
	Name        unsafe.Pointer
}

type _vkExportMemoryWin32HandleInfoKHR struct {
	sType       StructureType
	pNext       unsafe.Pointer
	pAttributes *windows.SecurityAttributes
	dwAccess    uint32
	name        unsafe.Pointer
}

func (s *_vkExportMemoryWin32HandleInfoKHR) Goify() *ExportMemoryWin32HandleInfoKHR {
	rval := &ExportMemoryWin32HandleInfoKHR{
		PNext:       (unsafe.Pointer)(s.pNext),
		PAttributes: (*windows.SecurityAttributes)(s.pAttributes),
		DwAccess:    (uint32)(s.dwAccess),
		Name:        (unsafe.Pointer)(s.name),
	}
	return rval
}
func (s *ExportMemoryWin32HandleInfoKHR) Vulkanize() *_vkExportMemoryWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkExportMemoryWin32HandleInfoKHR{
		sType:       STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR, /*c1*/
		pNext:       (unsafe.Pointer)(s.PNext),                          /*cb*/
		pAttributes: (*windows.SecurityAttributes)(s.PAttributes),       /*cb*/
		dwAccess:    (uint32)(s.DwAccess),                               /*cb*/
		name:        (unsafe.Pointer)(s.Name),                           /*cb*/
	}
	return rval
}

// ExportMemoryWin32HandleInfoNV: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkExportMemoryWin32HandleInfoNV.html
type ExportMemoryWin32HandleInfoNV struct {
	// SType = STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV
	PNext       unsafe.Pointer
	PAttributes *windows.SecurityAttributes
	DwAccess    uint32
}

type _vkExportMemoryWin32HandleInfoNV struct {
	sType       StructureType
	pNext       unsafe.Pointer
	pAttributes *windows.SecurityAttributes
	dwAccess    uint32
}

func (s *_vkExportMemoryWin32HandleInfoNV) Goify() *ExportMemoryWin32HandleInfoNV {
	rval := &ExportMemoryWin32HandleInfoNV{
		PNext:       (unsafe.Pointer)(s.pNext),
		PAttributes: (*windows.SecurityAttributes)(s.pAttributes),
		DwAccess:    (uint32)(s.dwAccess),
	}
	return rval
}
func (s *ExportMemoryWin32HandleInfoNV) Vulkanize() *_vkExportMemoryWin32HandleInfoNV {
	if s == nil {
		return nil
	}
	rval := &_vkExportMemoryWin32HandleInfoNV{
		sType:       STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV, /*c1*/
		pNext:       (unsafe.Pointer)(s.PNext),                         /*cb*/
		pAttributes: (*windows.SecurityAttributes)(s.PAttributes),      /*cb*/
		dwAccess:    (uint32)(s.DwAccess),                              /*cb*/
	}
	return rval
}

// ExportSemaphoreWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreWin32HandleInfoKHR.html
type ExportSemaphoreWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR
	PNext       unsafe.Pointer
	PAttributes *windows.SecurityAttributes
	DwAccess    uint32
	Name        unsafe.Pointer
}

type _vkExportSemaphoreWin32HandleInfoKHR struct {
	sType       StructureType
	pNext       unsafe.Pointer
	pAttributes *windows.SecurityAttributes
	dwAccess    uint32
	name        unsafe.Pointer
}

func (s *_vkExportSemaphoreWin32HandleInfoKHR) Goify() *ExportSemaphoreWin32HandleInfoKHR {
	rval := &ExportSemaphoreWin32HandleInfoKHR{
		PNext:       (unsafe.Pointer)(s.pNext),
		PAttributes: (*windows.SecurityAttributes)(s.pAttributes),
		DwAccess:    (uint32)(s.dwAccess),
		Name:        (unsafe.Pointer)(s.name),
	}
	return rval
}
func (s *ExportSemaphoreWin32HandleInfoKHR) Vulkanize() *_vkExportSemaphoreWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkExportSemaphoreWin32HandleInfoKHR{
		sType:       STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR, /*c1*/
		pNext:       (unsafe.Pointer)(s.PNext),                             /*cb*/
		pAttributes: (*windows.SecurityAttributes)(s.PAttributes),          /*cb*/
		dwAccess:    (uint32)(s.DwAccess),                                  /*cb*/
		name:        (unsafe.Pointer)(s.Name),                              /*cb*/
	}
	return rval
}

// FenceGetWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFenceGetWin32HandleInfoKHR.html
type FenceGetWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR
	PNext      unsafe.Pointer
	Fence      Fence
	HandleType ExternalFenceHandleTypeFlagBits
}

type _vkFenceGetWin32HandleInfoKHR struct {
	sType      StructureType
	pNext      unsafe.Pointer
	fence      Fence
	handleType ExternalFenceHandleTypeFlagBits
}

func (s *_vkFenceGetWin32HandleInfoKHR) Goify() *FenceGetWin32HandleInfoKHR {
	rval := &FenceGetWin32HandleInfoKHR{
		PNext:      (unsafe.Pointer)(s.pNext),
		Fence:      (Fence)(s.fence),
		HandleType: (ExternalFenceHandleTypeFlagBits)(s.handleType),
	}
	return rval
}
func (s *FenceGetWin32HandleInfoKHR) Vulkanize() *_vkFenceGetWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkFenceGetWin32HandleInfoKHR{
		sType:      STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR,  /*c1*/
		pNext:      (unsafe.Pointer)(s.PNext),                       /*cb*/
		fence:      (Fence)(s.Fence),                                /*cb*/
		handleType: (ExternalFenceHandleTypeFlagBits)(s.HandleType), /*cb*/
	}
	return rval
}

// ImportFenceWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html
type ImportFenceWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR
	PNext      unsafe.Pointer
	Fence      Fence
	Flags      FenceImportFlags
	HandleType ExternalFenceHandleTypeFlagBits
	Handle     windows.Handle
	Name       unsafe.Pointer
}

type _vkImportFenceWin32HandleInfoKHR struct {
	sType      StructureType
	pNext      unsafe.Pointer
	fence      Fence
	flags      FenceImportFlags
	handleType ExternalFenceHandleTypeFlagBits
	handle     windows.Handle
	name       unsafe.Pointer
}

func (s *_vkImportFenceWin32HandleInfoKHR) Goify() *ImportFenceWin32HandleInfoKHR {
	rval := &ImportFenceWin32HandleInfoKHR{
		PNext:      (unsafe.Pointer)(s.pNext),
		Fence:      (Fence)(s.fence),
		Flags:      (FenceImportFlags)(s.flags),
		HandleType: (ExternalFenceHandleTypeFlagBits)(s.handleType),
		Handle:     (windows.Handle)(s.handle),
		Name:       (unsafe.Pointer)(s.name),
	}
	return rval
}
func (s *ImportFenceWin32HandleInfoKHR) Vulkanize() *_vkImportFenceWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkImportFenceWin32HandleInfoKHR{
		sType:      STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR, /*c1*/
		pNext:      (unsafe.Pointer)(s.PNext),                         /*cb*/
		fence:      (Fence)(s.Fence),                                  /*cb*/
		flags:      (FenceImportFlags)(s.Flags),                       /*cb*/
		handleType: (ExternalFenceHandleTypeFlagBits)(s.HandleType),   /*cb*/
		handle:     (windows.Handle)(s.Handle),                        /*cb*/
		name:       (unsafe.Pointer)(s.Name),                          /*cb*/
	}
	return rval
}

// ImportMemoryWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoKHR.html
type ImportMemoryWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR
	PNext      unsafe.Pointer
	HandleType ExternalMemoryHandleTypeFlagBits
	Handle     windows.Handle
	Name       unsafe.Pointer
}

type _vkImportMemoryWin32HandleInfoKHR struct {
	sType      StructureType
	pNext      unsafe.Pointer
	handleType ExternalMemoryHandleTypeFlagBits
	handle     windows.Handle
	name       unsafe.Pointer
}

func (s *_vkImportMemoryWin32HandleInfoKHR) Goify() *ImportMemoryWin32HandleInfoKHR {
	rval := &ImportMemoryWin32HandleInfoKHR{
		PNext:      (unsafe.Pointer)(s.pNext),
		HandleType: (ExternalMemoryHandleTypeFlagBits)(s.handleType),
		Handle:     (windows.Handle)(s.handle),
		Name:       (unsafe.Pointer)(s.name),
	}
	return rval
}
func (s *ImportMemoryWin32HandleInfoKHR) Vulkanize() *_vkImportMemoryWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkImportMemoryWin32HandleInfoKHR{
		sType:      STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR, /*c1*/
		pNext:      (unsafe.Pointer)(s.PNext),                          /*cb*/
		handleType: (ExternalMemoryHandleTypeFlagBits)(s.HandleType),   /*cb*/
		handle:     (windows.Handle)(s.Handle),                         /*cb*/
		name:       (unsafe.Pointer)(s.Name),                           /*cb*/
	}
	return rval
}

// ImportMemoryWin32HandleInfoNV: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoNV.html
type ImportMemoryWin32HandleInfoNV struct {
	// SType = STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV
	PNext      unsafe.Pointer
	HandleType ExternalMemoryHandleTypeFlagsNV
	Handle     windows.Handle
}

type _vkImportMemoryWin32HandleInfoNV struct {
	sType      StructureType
	pNext      unsafe.Pointer
	handleType ExternalMemoryHandleTypeFlagsNV
	handle     windows.Handle
}

func (s *_vkImportMemoryWin32HandleInfoNV) Goify() *ImportMemoryWin32HandleInfoNV {
	rval := &ImportMemoryWin32HandleInfoNV{
		PNext:      (unsafe.Pointer)(s.pNext),
		HandleType: (ExternalMemoryHandleTypeFlagsNV)(s.handleType),
		Handle:     (windows.Handle)(s.handle),
	}
	return rval
}
func (s *ImportMemoryWin32HandleInfoNV) Vulkanize() *_vkImportMemoryWin32HandleInfoNV {
	if s == nil {
		return nil
	}
	rval := &_vkImportMemoryWin32HandleInfoNV{
		sType:      STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV, /*c1*/
		pNext:      (unsafe.Pointer)(s.PNext),                         /*cb*/
		handleType: (ExternalMemoryHandleTypeFlagsNV)(s.HandleType),   /*cb*/
		handle:     (windows.Handle)(s.Handle),                        /*cb*/
	}
	return rval
}

// ImportSemaphoreWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html
type ImportSemaphoreWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR
	PNext      unsafe.Pointer
	Semaphore  Semaphore
	Flags      SemaphoreImportFlags
	HandleType ExternalSemaphoreHandleTypeFlagBits
	Handle     windows.Handle
	Name       unsafe.Pointer
}

type _vkImportSemaphoreWin32HandleInfoKHR struct {
	sType      StructureType
	pNext      unsafe.Pointer
	semaphore  Semaphore
	flags      SemaphoreImportFlags
	handleType ExternalSemaphoreHandleTypeFlagBits
	handle     windows.Handle
	name       unsafe.Pointer
}

func (s *_vkImportSemaphoreWin32HandleInfoKHR) Goify() *ImportSemaphoreWin32HandleInfoKHR {
	rval := &ImportSemaphoreWin32HandleInfoKHR{
		PNext:      (unsafe.Pointer)(s.pNext),
		Semaphore:  (Semaphore)(s.semaphore),
		Flags:      (SemaphoreImportFlags)(s.flags),
		HandleType: (ExternalSemaphoreHandleTypeFlagBits)(s.handleType),
		Handle:     (windows.Handle)(s.handle),
		Name:       (unsafe.Pointer)(s.name),
	}
	return rval
}
func (s *ImportSemaphoreWin32HandleInfoKHR) Vulkanize() *_vkImportSemaphoreWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkImportSemaphoreWin32HandleInfoKHR{
		sType:      STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR, /*c1*/
		pNext:      (unsafe.Pointer)(s.PNext),                             /*cb*/
		semaphore:  (Semaphore)(s.Semaphore),                              /*cb*/
		flags:      (SemaphoreImportFlags)(s.Flags),                       /*cb*/
		handleType: (ExternalSemaphoreHandleTypeFlagBits)(s.HandleType),   /*cb*/
		handle:     (windows.Handle)(s.Handle),                            /*cb*/
		name:       (unsafe.Pointer)(s.Name),                              /*cb*/
	}
	return rval
}

// MemoryGetWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkMemoryGetWin32HandleInfoKHR.html
type MemoryGetWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR
	PNext      unsafe.Pointer
	Memory     DeviceMemory
	HandleType ExternalMemoryHandleTypeFlagBits
}

type _vkMemoryGetWin32HandleInfoKHR struct {
	sType      StructureType
	pNext      unsafe.Pointer
	memory     DeviceMemory
	handleType ExternalMemoryHandleTypeFlagBits
}

func (s *_vkMemoryGetWin32HandleInfoKHR) Goify() *MemoryGetWin32HandleInfoKHR {
	rval := &MemoryGetWin32HandleInfoKHR{
		PNext:      (unsafe.Pointer)(s.pNext),
		Memory:     (DeviceMemory)(s.memory),
		HandleType: (ExternalMemoryHandleTypeFlagBits)(s.handleType),
	}
	return rval
}
func (s *MemoryGetWin32HandleInfoKHR) Vulkanize() *_vkMemoryGetWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkMemoryGetWin32HandleInfoKHR{
		sType:      STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR,  /*c1*/
		pNext:      (unsafe.Pointer)(s.PNext),                        /*cb*/
		memory:     (DeviceMemory)(s.Memory),                         /*cb*/
		handleType: (ExternalMemoryHandleTypeFlagBits)(s.HandleType), /*cb*/
	}
	return rval
}

// MemoryWin32HandlePropertiesKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkMemoryWin32HandlePropertiesKHR.html
type MemoryWin32HandlePropertiesKHR struct {
	// SType = STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

// WARNING - struct MemoryWin32HandlePropertiesKHR is returned only, which is not yet handled in the binding
type _vkMemoryWin32HandlePropertiesKHR struct {
	sType          StructureType
	pNext          unsafe.Pointer
	memoryTypeBits uint32
}

func (s *_vkMemoryWin32HandlePropertiesKHR) Goify() *MemoryWin32HandlePropertiesKHR {
	rval := &MemoryWin32HandlePropertiesKHR{
		PNext:          (unsafe.Pointer)(s.pNext),
		MemoryTypeBits: (uint32)(s.memoryTypeBits),
	}
	return rval
}
func (s *MemoryWin32HandlePropertiesKHR) Vulkanize() *_vkMemoryWin32HandlePropertiesKHR {
	if s == nil {
		return nil
	}
	rval := &_vkMemoryWin32HandlePropertiesKHR{
		sType:          STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR, /*c1*/
		pNext:          (unsafe.Pointer)(s.PNext),                         /*cb*/
		memoryTypeBits: (uint32)(s.MemoryTypeBits),                        /*cb*/
	}
	return rval
}

// PhysicalDeviceSurfaceInfo2KHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html
type PhysicalDeviceSurfaceInfo2KHR struct {
	// SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR
	PNext   unsafe.Pointer
	Surface SurfaceKHR
}

type _vkPhysicalDeviceSurfaceInfo2KHR struct {
	sType   StructureType
	pNext   unsafe.Pointer
	surface SurfaceKHR
}

func (s *_vkPhysicalDeviceSurfaceInfo2KHR) Goify() *PhysicalDeviceSurfaceInfo2KHR {
	rval := &PhysicalDeviceSurfaceInfo2KHR{
		PNext:   (unsafe.Pointer)(s.pNext),
		Surface: (SurfaceKHR)(s.surface),
	}
	return rval
}
func (s *PhysicalDeviceSurfaceInfo2KHR) Vulkanize() *_vkPhysicalDeviceSurfaceInfo2KHR {
	if s == nil {
		return nil
	}
	rval := &_vkPhysicalDeviceSurfaceInfo2KHR{
		sType:   STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR, /*c1*/
		pNext:   (unsafe.Pointer)(s.PNext),                         /*cb*/
		surface: (SurfaceKHR)(s.Surface),                           /*cb*/
	}
	return rval
}

// SemaphoreGetWin32HandleInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSemaphoreGetWin32HandleInfoKHR.html
type SemaphoreGetWin32HandleInfoKHR struct {
	// SType = STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR
	PNext      unsafe.Pointer
	Semaphore  Semaphore
	HandleType ExternalSemaphoreHandleTypeFlagBits
}

type _vkSemaphoreGetWin32HandleInfoKHR struct {
	sType      StructureType
	pNext      unsafe.Pointer
	semaphore  Semaphore
	handleType ExternalSemaphoreHandleTypeFlagBits
}

func (s *_vkSemaphoreGetWin32HandleInfoKHR) Goify() *SemaphoreGetWin32HandleInfoKHR {
	rval := &SemaphoreGetWin32HandleInfoKHR{
		PNext:      (unsafe.Pointer)(s.pNext),
		Semaphore:  (Semaphore)(s.semaphore),
		HandleType: (ExternalSemaphoreHandleTypeFlagBits)(s.handleType),
	}
	return rval
}
func (s *SemaphoreGetWin32HandleInfoKHR) Vulkanize() *_vkSemaphoreGetWin32HandleInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkSemaphoreGetWin32HandleInfoKHR{
		sType:      STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR,  /*c1*/
		pNext:      (unsafe.Pointer)(s.PNext),                           /*cb*/
		semaphore:  (Semaphore)(s.Semaphore),                            /*cb*/
		handleType: (ExternalSemaphoreHandleTypeFlagBits)(s.HandleType), /*cb*/
	}
	return rval
}

// SurfaceCapabilitiesFullScreenExclusiveEXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesFullScreenExclusiveEXT.html
type SurfaceCapabilitiesFullScreenExclusiveEXT struct {
	// SType = STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT
	PNext                        unsafe.Pointer
	FullScreenExclusiveSupported bool
}

type _vkSurfaceCapabilitiesFullScreenExclusiveEXT struct {
	sType                        StructureType
	pNext                        unsafe.Pointer
	fullScreenExclusiveSupported Bool32
}

func (s *_vkSurfaceCapabilitiesFullScreenExclusiveEXT) Goify() *SurfaceCapabilitiesFullScreenExclusiveEXT {
	rval := &SurfaceCapabilitiesFullScreenExclusiveEXT{
		PNext:                        (unsafe.Pointer)(s.pNext),
		FullScreenExclusiveSupported: translatePublic_Bool32(s.fullScreenExclusiveSupported), /*default*/
	}
	return rval
}
func (s *SurfaceCapabilitiesFullScreenExclusiveEXT) Vulkanize() *_vkSurfaceCapabilitiesFullScreenExclusiveEXT {
	if s == nil {
		return nil
	}
	rval := &_vkSurfaceCapabilitiesFullScreenExclusiveEXT{
		sType:                        STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT, /*c1*/
		pNext:                        (unsafe.Pointer)(s.PNext),                                     /*cb*/
		fullScreenExclusiveSupported: translateInternal_Bool32(s.FullScreenExclusiveSupported),      /*default*/
	}
	return rval
}

// SurfaceFullScreenExclusiveInfoEXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html
type SurfaceFullScreenExclusiveInfoEXT struct {
	// SType = STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT
	PNext               unsafe.Pointer
	FullScreenExclusive FullScreenExclusiveEXT
}

type _vkSurfaceFullScreenExclusiveInfoEXT struct {
	sType               StructureType
	pNext               unsafe.Pointer
	fullScreenExclusive FullScreenExclusiveEXT
}

func (s *_vkSurfaceFullScreenExclusiveInfoEXT) Goify() *SurfaceFullScreenExclusiveInfoEXT {
	rval := &SurfaceFullScreenExclusiveInfoEXT{
		PNext:               (unsafe.Pointer)(s.pNext),
		FullScreenExclusive: (FullScreenExclusiveEXT)(s.fullScreenExclusive),
	}
	return rval
}
func (s *SurfaceFullScreenExclusiveInfoEXT) Vulkanize() *_vkSurfaceFullScreenExclusiveInfoEXT {
	if s == nil {
		return nil
	}
	rval := &_vkSurfaceFullScreenExclusiveInfoEXT{
		sType:               STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT, /*c1*/
		pNext:               (unsafe.Pointer)(s.PNext),                             /*cb*/
		fullScreenExclusive: (FullScreenExclusiveEXT)(s.FullScreenExclusive),       /*cb*/
	}
	return rval
}

// SurfaceFullScreenExclusiveWin32InfoEXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html
type SurfaceFullScreenExclusiveWin32InfoEXT struct {
	// SType = STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT
	PNext    unsafe.Pointer
	Hmonitor windows.Handle
}

type _vkSurfaceFullScreenExclusiveWin32InfoEXT struct {
	sType    StructureType
	pNext    unsafe.Pointer
	hmonitor windows.Handle
}

func (s *_vkSurfaceFullScreenExclusiveWin32InfoEXT) Goify() *SurfaceFullScreenExclusiveWin32InfoEXT {
	rval := &SurfaceFullScreenExclusiveWin32InfoEXT{
		PNext:    (unsafe.Pointer)(s.pNext),
		Hmonitor: (windows.Handle)(s.hmonitor),
	}
	return rval
}
func (s *SurfaceFullScreenExclusiveWin32InfoEXT) Vulkanize() *_vkSurfaceFullScreenExclusiveWin32InfoEXT {
	if s == nil {
		return nil
	}
	rval := &_vkSurfaceFullScreenExclusiveWin32InfoEXT{
		sType:    STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT, /*c1*/
		pNext:    (unsafe.Pointer)(s.PNext),                                   /*cb*/
		hmonitor: (windows.Handle)(s.Hmonitor),                                /*cb*/
	}
	return rval
}

// Win32KeyedMutexAcquireReleaseInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html
type Win32KeyedMutexAcquireReleaseInfoKHR struct {
	// SType = STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR
	PNext unsafe.Pointer
	// acquireCount
	PAcquireSyncs    []DeviceMemory
	PAcquireKeys     []uint64
	PAcquireTimeouts []uint32
	// releaseCount
	PReleaseSyncs []DeviceMemory
	PReleaseKeys  []uint64
}

type _vkWin32KeyedMutexAcquireReleaseInfoKHR struct {
	sType            StructureType
	pNext            unsafe.Pointer
	acquireCount     uint32
	pAcquireSyncs    *DeviceMemory
	pAcquireKeys     *uint64
	pAcquireTimeouts *uint32
	releaseCount     uint32
	pReleaseSyncs    *DeviceMemory
	pReleaseKeys     *uint64
}

func (s *_vkWin32KeyedMutexAcquireReleaseInfoKHR) Goify() *Win32KeyedMutexAcquireReleaseInfoKHR {
	rval := &Win32KeyedMutexAcquireReleaseInfoKHR{
		PNext: (unsafe.Pointer)(s.pNext),
		// Unexpected 'isLenForAnotherMember'!
		// Unexpected 'isLenForAnotherMember'!
	}
	return rval
}
func (s *Win32KeyedMutexAcquireReleaseInfoKHR) Vulkanize() *_vkWin32KeyedMutexAcquireReleaseInfoKHR {
	if s == nil {
		return nil
	}

	var psl_pAcquireSyncs *DeviceMemory
	if len(s.PAcquireSyncs) > 0 {
		psl_pAcquireSyncs = &s.PAcquireSyncs[0]
	}

	var psl_pAcquireKeys *uint64
	if len(s.PAcquireKeys) > 0 {
		psl_pAcquireKeys = &s.PAcquireKeys[0]
	}

	var psl_pAcquireTimeouts *uint32
	if len(s.PAcquireTimeouts) > 0 {
		psl_pAcquireTimeouts = &s.PAcquireTimeouts[0]
	}

	var psl_pReleaseSyncs *DeviceMemory
	if len(s.PReleaseSyncs) > 0 {
		psl_pReleaseSyncs = &s.PReleaseSyncs[0]
	}

	var psl_pReleaseKeys *uint64
	if len(s.PReleaseKeys) > 0 {
		psl_pReleaseKeys = &s.PReleaseKeys[0]
	}
	rval := &_vkWin32KeyedMutexAcquireReleaseInfoKHR{
		sType:            STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR, /*c1*/
		pNext:            (unsafe.Pointer)(s.PNext),                                 /*cb*/
		pAcquireSyncs:    psl_pAcquireSyncs,                                         /*c rem*/
		pAcquireKeys:     psl_pAcquireKeys,                                          /*c rem*/
		pAcquireTimeouts: psl_pAcquireTimeouts,                                      /*c rem*/
		pReleaseSyncs:    psl_pReleaseSyncs,                                         /*c rem*/
		pReleaseKeys:     psl_pReleaseKeys,                                          /*c rem*/
	}
	rval.acquireCount = 0 // c6-b
	if uint32(len(s.PAcquireSyncs)) > rval.acquireCount {
		rval.acquireCount = uint32(len(s.PAcquireSyncs))
	}
	if uint32(len(s.PAcquireKeys)) > rval.acquireCount {
		rval.acquireCount = uint32(len(s.PAcquireKeys))
	}
	if uint32(len(s.PAcquireTimeouts)) > rval.acquireCount {
		rval.acquireCount = uint32(len(s.PAcquireTimeouts))
	}
	rval.releaseCount = 0 // c6-b
	if uint32(len(s.PReleaseSyncs)) > rval.releaseCount {
		rval.releaseCount = uint32(len(s.PReleaseSyncs))
	}
	if uint32(len(s.PReleaseKeys)) > rval.releaseCount {
		rval.releaseCount = uint32(len(s.PReleaseKeys))
	}
	return rval
}

// Win32KeyedMutexAcquireReleaseInfoNV: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html
type Win32KeyedMutexAcquireReleaseInfoNV struct {
	// SType = STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV
	PNext unsafe.Pointer
	// acquireCount
	PAcquireSyncs               []DeviceMemory
	PAcquireKeys                []uint64
	PAcquireTimeoutMilliseconds []uint32
	// releaseCount
	PReleaseSyncs []DeviceMemory
	PReleaseKeys  []uint64
}

type _vkWin32KeyedMutexAcquireReleaseInfoNV struct {
	sType                       StructureType
	pNext                       unsafe.Pointer
	acquireCount                uint32
	pAcquireSyncs               *DeviceMemory
	pAcquireKeys                *uint64
	pAcquireTimeoutMilliseconds *uint32
	releaseCount                uint32
	pReleaseSyncs               *DeviceMemory
	pReleaseKeys                *uint64
}

func (s *_vkWin32KeyedMutexAcquireReleaseInfoNV) Goify() *Win32KeyedMutexAcquireReleaseInfoNV {
	rval := &Win32KeyedMutexAcquireReleaseInfoNV{
		PNext: (unsafe.Pointer)(s.pNext),
		// Unexpected 'isLenForAnotherMember'!
		// Unexpected 'isLenForAnotherMember'!
	}
	return rval
}
func (s *Win32KeyedMutexAcquireReleaseInfoNV) Vulkanize() *_vkWin32KeyedMutexAcquireReleaseInfoNV {
	if s == nil {
		return nil
	}

	var psl_pAcquireSyncs *DeviceMemory
	if len(s.PAcquireSyncs) > 0 {
		psl_pAcquireSyncs = &s.PAcquireSyncs[0]
	}

	var psl_pAcquireKeys *uint64
	if len(s.PAcquireKeys) > 0 {
		psl_pAcquireKeys = &s.PAcquireKeys[0]
	}

	var psl_pAcquireTimeoutMilliseconds *uint32
	if len(s.PAcquireTimeoutMilliseconds) > 0 {
		psl_pAcquireTimeoutMilliseconds = &s.PAcquireTimeoutMilliseconds[0]
	}

	var psl_pReleaseSyncs *DeviceMemory
	if len(s.PReleaseSyncs) > 0 {
		psl_pReleaseSyncs = &s.PReleaseSyncs[0]
	}

	var psl_pReleaseKeys *uint64
	if len(s.PReleaseKeys) > 0 {
		psl_pReleaseKeys = &s.PReleaseKeys[0]
	}
	rval := &_vkWin32KeyedMutexAcquireReleaseInfoNV{
		sType:                       STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV, /*c1*/
		pNext:                       (unsafe.Pointer)(s.PNext),                                /*cb*/
		pAcquireSyncs:               psl_pAcquireSyncs,                                        /*c rem*/
		pAcquireKeys:                psl_pAcquireKeys,                                         /*c rem*/
		pAcquireTimeoutMilliseconds: psl_pAcquireTimeoutMilliseconds,                          /*c rem*/
		pReleaseSyncs:               psl_pReleaseSyncs,                                        /*c rem*/
		pReleaseKeys:                psl_pReleaseKeys,                                         /*c rem*/
	}
	rval.acquireCount = 0 // c6-b
	if uint32(len(s.PAcquireSyncs)) > rval.acquireCount {
		rval.acquireCount = uint32(len(s.PAcquireSyncs))
	}
	if uint32(len(s.PAcquireKeys)) > rval.acquireCount {
		rval.acquireCount = uint32(len(s.PAcquireKeys))
	}
	if uint32(len(s.PAcquireTimeoutMilliseconds)) > rval.acquireCount {
		rval.acquireCount = uint32(len(s.PAcquireTimeoutMilliseconds))
	}
	rval.releaseCount = 0 // c6-b
	if uint32(len(s.PReleaseSyncs)) > rval.releaseCount {
		rval.releaseCount = uint32(len(s.PReleaseSyncs))
	}
	if uint32(len(s.PReleaseKeys)) > rval.releaseCount {
		rval.releaseCount = uint32(len(s.PReleaseKeys))
	}
	return rval
}

// Win32SurfaceCreateInfoKHR: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkWin32SurfaceCreateInfoKHR.html
type Win32SurfaceCreateInfoKHR struct {
	// SType = STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR
	PNext     unsafe.Pointer
	Flags     Win32SurfaceCreateFlagsKHR
	Hinstance windows.Handle
	Hwnd      windows.HWND
}

type _vkWin32SurfaceCreateInfoKHR struct {
	sType     StructureType
	pNext     unsafe.Pointer
	flags     Win32SurfaceCreateFlagsKHR
	hinstance windows.Handle
	hwnd      windows.HWND
}

func (s *_vkWin32SurfaceCreateInfoKHR) Goify() *Win32SurfaceCreateInfoKHR {
	rval := &Win32SurfaceCreateInfoKHR{
		PNext:     (unsafe.Pointer)(s.pNext),
		Flags:     (Win32SurfaceCreateFlagsKHR)(s.flags),
		Hinstance: (windows.Handle)(s.hinstance),
		Hwnd:      (windows.HWND)(s.hwnd),
	}
	return rval
}
func (s *Win32SurfaceCreateInfoKHR) Vulkanize() *_vkWin32SurfaceCreateInfoKHR {
	if s == nil {
		return nil
	}
	rval := &_vkWin32SurfaceCreateInfoKHR{
		sType:     STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR, /*c1*/
		pNext:     (unsafe.Pointer)(s.PNext),                    /*cb*/
		flags:     (Win32SurfaceCreateFlagsKHR)(s.Flags),        /*cb*/
		hinstance: (windows.Handle)(s.Hinstance),                /*cb*/
		hwnd:      (windows.HWND)(s.Hwnd),                       /*cb*/
	}
	return rval
}
