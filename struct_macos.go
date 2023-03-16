//go:build darwin
// Code generated by go-vk from vk-1.2.203.xml at 2023-03-16 13:28:32.599214 -0600 MDT m=+2.629418270. DO NOT EDIT.

package vk

import "unsafe"

// MacOSSurfaceCreateInfoMVK: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkMacOSSurfaceCreateInfoMVK.html
type MacOSSurfaceCreateInfoMVK struct {
	// SType = STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK
	PNext unsafe.Pointer
	Flags MacOSSurfaceCreateFlagsMVK
	PView unsafe.Pointer
}

type _vkMacOSSurfaceCreateInfoMVK struct {
	sType StructureType
	pNext unsafe.Pointer
	flags MacOSSurfaceCreateFlagsMVK
	pView unsafe.Pointer
}

func (s *_vkMacOSSurfaceCreateInfoMVK) Goify() *MacOSSurfaceCreateInfoMVK {
	rval := &MacOSSurfaceCreateInfoMVK{
		PNext: (unsafe.Pointer)(s.pNext),
		Flags: (MacOSSurfaceCreateFlagsMVK)(s.flags),
		PView: (unsafe.Pointer)(s.pView),
	}
	return rval
}
func (s *MacOSSurfaceCreateInfoMVK) Vulkanize() *_vkMacOSSurfaceCreateInfoMVK {
	if s == nil {
		return nil
	}
	rval := &_vkMacOSSurfaceCreateInfoMVK{
		sType: STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK, /*c1*/
		pNext: (unsafe.Pointer)(s.PNext),                    /*cb*/
		flags: (MacOSSurfaceCreateFlagsMVK)(s.Flags),        /*cb*/
		pView: (unsafe.Pointer)(s.PView),                    /*cb*/
	}
	return rval
}
