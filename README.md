# go-vk

go-vk is a Go-langauge (and Go-style) binding around the Vulkan API. Rather than just slapping a Cgo wrapper around
everything, Vulkan's functions, structures and other types have been translated to a Go-style API.

For example, "native" Vulkan returns any resources you request in pointers your program passes into Vulkan. This
allows Vulkan to always return a VkResult success or error code from the C function call. However, in Go,
we have the luxury of multiple return values, so this:

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

Likewise, functions returning an array of values in C require a call, an error check, an allocation, 
another function call, and another error check:
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
// Now do something with the devices and make sure you hold on to deviceCount so you don't go beyond the bounds of the array...
```

Yuck. Here's the same code in Go:
```go
devices, err := vk.EnumeratePhysicalDevices(myInstance)
if err != nil {
    // handle the error
}
// devices is a slice of vk.PhysicalDevice. Nice!
```

This codebase is (almost) entirely generated from a `vk.xml` file by the `[vk-gen](https://github.com/bbredesen/vk-gen)` toolset. Updating go-vk for a new 
Vulkan version should be as easy as downloading the new vk.xml file from Khronos and executing vk-gen.

# Usage

Ensure that your GPU supports Vulkan and that a Vulkan library is installed in your system-default 
library location (e.g., C:\windows\system32\vulkan-1.dll on Windows).

`go get go-vk@latest`

```go
package main

import (
    "github.com/bbredesen/go-vk"
)

// Notice that you don't need to alias the import, it is already bound to the "vk" namespace

func main() {
    if encodedVersion, err := vk.EnumerateInstanceVersion(); err != nil {
        fmt.Printf("EnumerateInstanceVersion failed! Error code was %s\n", err.Error())
    } else {
        fmt.Printf("Installed Vulkan version: %d.%d.%d\n", 
            vk.API_VERSION_MAJOR(encodedVersion), 
            vk.API_VERSION_MINOR(encodedVersion), 
            vk.API_VERSION_PATCH(encodedVersion),
        )
    }
}
```

`go run main.go`

# Library Structure

The Vulkan API is defined with a set of type categories, each of which has a corresponding source file in go-vk. Thus, you will find
all structs defined in struct.go, all bitmask types and values defined in bitmask.go, etc. Where platform-specific
types are neccessary, they are defined in separate files with appropriate go:build tags. The `stringify` tool has also been run against enumerated types.

The underlying Vulkan implementation is actually accessed through a small Cgo wrapper, found in static_common.go; go-vk opens the shared library
and lazy-loads any requested symbols. All of the public-facing structs in Go are translated to the appropriate memory
layout before being passed through to the API, via each struct's Vulkanize() function. Vulkanize()'s primary
purpose is to convert slices to a length and pointer field in the internal struct, Go strings to null-terminated
byte pointers, and to recursively Vulkanize any non-primitive members.

A small number of structs that are returned by Vulkan calls also have a Goify function, doing the reverse: 
create slices from a length and pointer field and create strings from null-terminated byte arrays.

Note that you should never need to directly call Vulkanize() or Goify(). Conversions are automatically
handled in the background when you call a Vulkan function.

# Examples

See the `[go-vk-samples](https://github.com/bbredesen/go-vk-samples) repo for a number of working Vulkan samples using this library.
For the most part, the samples only run on Windows. 

Minimal testing of `go-vk` has been done against Macs/MoltenVK. No testing has been done on Linux or other platforms.