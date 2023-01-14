# go-vk

go-vk is a Go-langauge (and Go-style) binding around the Vulkan graphichs API. Rather than just slapping a Cgo wrapper
around everything, Vulkan's functions, structures and other types have been translated to a Go-style API.

For example, "native" Vulkan returns any resources you request in pointers your program passes into Vulkan. This allows
Vulkan to (almost) always return a VkResult success or error code from the C function call. However, in Go, we have the
luxury of multiple return values, so this:

```C
VkInstance myInstance;
Result res = vkCreateInstance(&instanceCI, NULL, &myInstance);
if (res != VK_SUCCESS) {
    // Handle an error
}
// Use the instance handle
```

Becomes this:
```go
myInstance, err := vk.CreateInstance(instanceCI, nil)
if err != nil { // VK_SUCCESS is returned as nil
    panic(err) // Don't panic
}
```

Likewise, the "Enumerate" group of functions returning an array of values in C require a call, an error check, an
allocation, another function call, and another error check:
```C
int deviceCount;
Result res = vkEnumeratePhysicalDevices(myInstance, &deviceCount, NULL);
if (res != VK_SUCCESS) { // Check the result, of course
    // handle the error
}
// ...and probably also check that deviceCount > 0

VkPhysicalDevice *devices = malloc(deviceCount * sizeof(VkPhysicalDevice));
res = vkEnumeratePhysicalDevices(myInstance, &deviceCount, devices);
if (res != VK_SUCCESS) { // Check the result again
    // handle the error
}
// Now do something with the devices and make sure you hold on to deviceCount 
// so you don't go beyond the bounds of the array...
```

Yuck. Here's the same code in Go:
```go
devices, err := vk.EnumeratePhysicalDevices(myInstance)
if err != nil {
    // handle the error
}
// devices is a slice of vk.PhysicalDevice. Nice!
```

This codebase is (almost) entirely generated from a `vk.xml` file by the `[vk-gen](https://github.com/bbredesen/vk-gen)`
tool. Updating go-vk for a new Vulkan version should be as easy as downloading the new vk.xml file from Khronos and
executing vk-gen. **This repo does not get direct modifications!** Any bug fixes or new features need to be made in
`vk-gen`, which will then re-generate this code base.

# Usage

Ensure that your GPU supports Vulkan and that a Vulkan library is installed in your system-default library location
(e.g., C:\windows\system32\vulkan-1.dll on Windows).

`$ go get go-vk@latest`

```go main.go
package main

import (
    "github.com/bbredesen/go-vk"
)
// Notice that you don't need to alias the import, it is already bound to the "vk" namespace

func main() {
    if r, encodedVersion := vk.EnumerateInstanceVersion(); r != vk.SUCCESS {
        fmt.Printf("EnumerateInstanceVersion failed! Error code was %s\n", err.Error())
        os.Exit(1)
    } else {
        fmt.Printf("Installed Vulkan version: %d.%d.%d\n", 
            vk.API_VERSION_MAJOR(encodedVersion), 
            vk.API_VERSION_MINOR(encodedVersion), 
            vk.API_VERSION_PATCH(encodedVersion),
        )
    }

    // Also notice that you don't need to set the StructureType field on your Go structs. 
    // In fact, it doesn't even exist on the public side of the binding...it is automatically
    // added when you pass your struct through to a command.
    appInfo := vk.ApplicationInfo{
		ApplicationName:    "Example App",
		ApplicationVersion: vk.MAKE_VERSION(1, 0, 0),
		EngineVersion:      vk.MAKE_VERSION(1, 0, 0),
		ApiVersion:         vk.MAKE_VERSION(1, 0, 0),
	}

	icInfo := vk.InstanceCreateInfo{
		ApplicationInfo:       appInfo,
		EnabledExtensionNames: app.requiredInstanceExt,
		EnabledLayerNames:     app.enableApiLayers,
	}

	r, instance := vk.CreateInstance(icInfo, nil)
    // r is actually a vk.Result, which is itself just an int32. It does 
    // implement both the the error interface and String(). Because it is 
    // returned as a value, not a pointer, you cannot test for nil (Go-style
    // error checking), but you are able to directly compare it to the known 
    // error codes that Vulkan might return.
    if r != vk.SUCCESS {
        fmt.Printf("Failed to create Vulkan instance, error code was %s\n", err.Error())
        if r == vk.ERROR_INCOMPATIBLE_DRIVER { 
            /* ... */
        }
    }
    fmt.Printf("Vulkan instance created, handle value is %X\n", instance)
}
```

`$ go run main.go`

A number of code samples and working demos are available at [go-vk-samples](https://github.com/bbredesen/go-vk-samples)


# Library Structure

The Vulkan API is defined through a set of type categories, each of which has a corresponding source file in go-vk.
Thus, you will find all structs defined in struct.go, all bitmask types defined in bitmask.go, etc. Where
platform-specific types are neccessary, they are defined in separate files with appropriate go:build tags. The
`stringify` tool has also been run against enumerated types, so if `result == vk.NOT_READY` then `result.String() == "NOT_READY"`.

The underlying Vulkan implementation is actually accessed through a small Cgo wrapper, found in static_common.go; go-vk
opens the shared library and lazy-loads any requested symbols. All of the public-facing structs in Go are translated to
the appropriate memory layout before being passed through to the API, via each struct's Vulkanize() function.
Vulkanize()'s primary purpose is to convert slices to a length and pointer field in the internal struct, Go strings to
null-terminated byte pointers, and to recursively Vulkanize any non-primitive members.

A small number of structs that are returned by Vulkan calls also have a Goify function, doing the reverse: create slices
from a length and pointer field and create strings from null-terminated byte arrays.

Note that you should never need to directly call Vulkanize() or Goify(). Conversions are automatically handled in the
background when you call a Vulkan function.

# Examples

See the `[go-vk-samples](https://github.com/bbredesen/go-vk-samples) repo for a number of working Vulkan samples using
this library. For the most part, the samples only run on Windows. 

Minimal testing of `go-vk` has been done against Macs/MoltenVK. Mac/MoltenVk versions of the samples will be coming in the future. No testing has been done on Linux or other platforms.