#ifndef __DLLOAD_H__
#define __DLLOAD_H__

#include <stdint.h>

void* OpenLibrary(const char *name);
void CloseLibrary(void *lib_handle);

void* SymbolFromName(void *lib_handle, const void *name);

// CallGetProcAddr calls a vkGetInstanceProcAddr function pointer. It exists so that command resolution does not depend on
// the generated trampolines (which themselves need resolved commands).
void* CallGetProcAddr(void *get_proc_addr, uintptr_t instance, const char *name);

#endif
