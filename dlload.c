
#include <stdio.h>
#include <stdint.h>
#include "dlload.h"

#ifndef _WIN32
#include <dlfcn.h>


void* OpenLibrary(const char *name) {
	void* lib_handle = dlopen(name, RTLD_LOCAL|RTLD_LAZY);
	if (!lib_handle) {
		printf("Unable to load %s: %s\n", name, dlerror());
		return NULL;
	}
	return lib_handle;
}

void CloseLibrary(void *lib_handle) {
	if (dlclose(lib_handle) != 0) {
		printf("Problem closing library: %s", dlerror());
	}
}

void* SymbolFromName(void *lib_handle, const void *name) {
	return dlsym(lib_handle, (const char*) name);
}


#endif
// typedef uint64_t (*vkCreateInstance_func)(const void*, const void*, uint64_t*);
typedef uint32_t (*vkEnumerateInstanceVersion_func)(uint32_t*);
typedef size_t (*vkGeneric_func3)(uintptr_t, uintptr_t, uintptr_t);
typedef size_t (*vkGeneric_func6)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef size_t (*vkGeneric_func9)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);


void* OpenLibrary(const char *name) { return 0; }
void CloseLibrary(void *lib_handle) {}
void* SymbolFromName(void *lib_handle, const void *name) { return 0; }

size_t Trampoline3(void *symbol, size_t p0, size_t p1, size_t p2) {
	return ((vkGeneric_func3) symbol)(p0, p1, p2);
}
size_t Trampoline6(void *symbol, size_t p0, size_t p1, size_t p2, size_t p3, size_t p4, size_t p5) {
	return ((vkGeneric_func6) symbol)(p0, p1, p2, p3, p4, p5);
}
size_t Trampoline9(void *symbol, size_t p0, size_t p1, size_t p2, size_t p3, size_t p4, size_t p5, size_t p6, size_t p7, size_t p8) {
	return ((vkGeneric_func9) symbol)(p0, p1, p2, p3, p4, p5, p6, p7, p8);
}

