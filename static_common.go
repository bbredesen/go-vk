package vk

import (
	"runtime"
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
// TODO: Enhance this method to provide error messages from the OS (?)
// TODO: A way for commands to indicate if the Result code is an error for that command, or an unexpected return value?
func (r Result) Error() string {
	return r.String()
}

type vkCommandKey int
type vkCommand struct {
	protoName string
	argCount  int
	hasReturn bool
	fnHandle  unsafe.Pointer
}

var lazyCommands map[vkCommandKey]vkCommand

var dlHandle unsafe.Pointer

func execTrampoline(commandKey vkCommandKey, args ...uintptr) uintptr {
	if dlHandle == nil {
		var libName string
		switch runtime.GOOS {
		case "windows":
			// This won't actually open the library on windows, since OpenLibrary is calling dlopen etc.
			// Need to either wrap dlopen, use a third party library, or generate different defs by platform.
			libName = "vulkan-1.dll"
		case "darwin":
			// TODO: Running on Mac/Darwin is tested only to the point of creating and
			// destroying a Vulkan instance.
			libName = "libMoltenVK.dylib"
		case "linux":
			// TODO: Opening/running on linux is completely untested.
			libName = "libvulkan.1.dylib"
		default:
			panic("Unsupported GOOS at OpenLibrary: " + runtime.GOOS)
		}
		cstr := C.CString(libName)
		dlHandle = C.OpenLibrary(cstr)
		C.free(unsafe.Pointer(cstr))
	}

	cmd := lazyCommands[commandKey]
	if cmd.fnHandle == nil {
		cmd.fnHandle = C.SymbolFromName(dlHandle, unsafe.Pointer(sys_stringToBytePointer(cmd.protoName)))
		lazyCommands[commandKey] = cmd
	}

	if len(args) != cmd.argCount {
		panic("Wrong number of arguments passed for cmd " + cmd.protoName)
	}

	var result C.uintptr_t
	// var err error

	switch cmd.argCount {
	case 1:
		result = C.Trampoline3(cmd.fnHandle, C.uintptr_t(args[0]), 0, 0)
	case 2:
		result = C.Trampoline3(cmd.fnHandle, C.uintptr_t(args[0]), C.uintptr_t(args[1]), 0)
	case 3:
		result = C.Trampoline3(cmd.fnHandle, C.uintptr_t(args[0]), C.uintptr_t(args[1]), C.uintptr_t(args[2]))
	case 4:
		result = C.Trampoline6(cmd.fnHandle, C.uintptr_t(args[0]), C.uintptr_t(args[1]), C.uintptr_t(args[2]), C.uintptr_t(args[3]), 0, 0)
	case 5:
		result = C.Trampoline6(cmd.fnHandle, C.uintptr_t(args[0]), C.uintptr_t(args[1]), C.uintptr_t(args[2]), C.uintptr_t(args[3]), C.uintptr_t(args[4]), 0)
	case 6:
		result = C.Trampoline6(cmd.fnHandle, C.uintptr_t(args[0]), C.uintptr_t(args[1]), C.uintptr_t(args[2]), C.uintptr_t(args[3]), C.uintptr_t(args[4]), C.uintptr_t(args[5]))
	case 7:
		result = C.Trampoline9(cmd.fnHandle, C.uintptr_t(args[0]), C.uintptr_t(args[1]), C.uintptr_t(args[2]), C.uintptr_t(args[3]), C.uintptr_t(args[4]), C.uintptr_t(args[5]), C.uintptr_t(args[6]), 0, 0)
	case 8:
		result = C.Trampoline9(cmd.fnHandle, C.uintptr_t(args[0]), C.uintptr_t(args[1]), C.uintptr_t(args[2]), C.uintptr_t(args[3]), C.uintptr_t(args[4]), C.uintptr_t(args[5]), C.uintptr_t(args[6]), C.uintptr_t(args[7]), 0)
	default:
		// There are no commands with 0 or 9+ arguments as of Vulkan 1.2.176
		panic("Unexpected number of arguments passed for cmd " + cmd.protoName)
	}

	// Trampoline is always returning a file does not exist error in the second return value, so that error reporting is disabled

	return uintptr(result) //, err
}

func stringToNullTermBytes(s string) *byte {
	b := []byte(s)
	b = append(b, 0)
	return &b[0]
}

// func getStringFromPtr(ptr uintptr, len int) string {
// 	var sl = struct {
// 		addr uintptr
// 		len  int
// 		cap  int
// 	}{ptr, len, len}
// 	return string(*(*[]byte)(unsafe.Pointer(&sl)))
// }
// func nullTermBytesToString(b *byte) string {
// 	rawPtr := uintptr(unsafe.Pointer(b))
// 	strLen := uintptr(0)

// 	var nextB byte
// 	nextB = *(*byte)(unsafe.Pointer(rawPtr + strLen))

// 	for nextB = *(*byte)(unsafe.Pointer(rawPtr)); nextB != 0; strLen++ {
// 	}

// 	return getStringFromPtr(rawPtr, int(strLen))
// }
