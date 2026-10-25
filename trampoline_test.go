package vk

import (
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// This test builds a small fake "Vulkan" library whose functions have the real C signatures of the commands that take
// floats by value, calls them through the generated bindings, and checks that the callee received the right values.
// The library records its arguments; they are read back through the same trampoline machinery using integer-only
// functions (float bits travel as integers).
//
// It needs a C compiler on PATH and is skipped otherwise (and always skipped on Windows). The test binary re-executes itself with GODEBUG settings that
// make memory-safety failures deterministic; see TestMain and escapes_test.go.

const fakeVulkanC = `
#include <stdint.h>
#include <string.h>
#include <unistd.h>

static float    g_f[4];
static uint64_t g_i[4];

void vkCmdSetLineWidth(uintptr_t cb, float w) { g_i[0] = cb; g_f[0] = w; }
void vkCmdSetDepthBias(uintptr_t cb, float a, float b, float c) { g_i[0] = cb; g_f[0] = a; g_f[1] = b; g_f[2] = c; }
void vkCmdSetDepthBounds(uintptr_t cb, float a, float b) { g_i[0] = cb; g_f[0] = a; g_f[1] = b; }
void vkSetDeviceMemoryPriorityEXT(uintptr_t dev, uint64_t mem, float p) { g_i[0] = dev; g_i[1] = mem; g_f[0] = p; }

// ---- support for escapes_test.go ----

// Commands that run "long" block for g_delay_us, as a real vkQueueSubmit or vkWaitForFences can.
static unsigned g_delay_us;
void testSetDelay(uintptr_t us) { g_delay_us = (unsigned)us; }

// 1 if the last vkQueueSubmit saw exactly the data the test built, -1 if anything was wrong
static int g_submit_ok;
uintptr_t testSubmitOK(uintptr_t unused) { return (uintptr_t)(intptr_t)g_submit_ok; }

// Mirrors VkSubmitInfo on 64 bit platforms
typedef struct {
	int32_t sType;
	const void *pNext;
	uint32_t waitSemaphoreCount;
	const uint64_t *pWaitSemaphores;
	const uint32_t *pWaitDstStageMask;
	uint32_t commandBufferCount;
	const uintptr_t *pCommandBuffers;
	uint32_t signalSemaphoreCount;
	const uint64_t *pSignalSemaphores;
} SubmitInfo;

// Every value in a submit is derivable from its position; escapes_test.go builds the same pattern.
int32_t vkQueueSubmit(uintptr_t queue, uint32_t count, const SubmitInfo *p, uint64_t fence) {
	usleep(g_delay_us);

	int ok = (queue == 0x51 && fence == 0xFE);
	for (uint32_t i = 0; ok && i < count; i++) {
		ok = p[i].sType == 4 /* STRUCTURE_TYPE_SUBMIT_INFO */ &&
			p[i].pNext != NULL && *(const uint64_t *)p[i].pNext == 0xFEEDFACE00000000ull + i &&
			p[i].waitSemaphoreCount == 64 && p[i].commandBufferCount == 64 && p[i].signalSemaphoreCount == 64;
		for (uint32_t j = 0; ok && j < 64; j++) {
			ok = p[i].pWaitSemaphores[j] == 0x1000 + i * 100 + j &&
				p[i].pWaitDstStageMask[j] == 0x2000 + i * 100 + j &&
				p[i].pCommandBuffers[j] == 0x3000 + i * 100 + j &&
				p[i].pSignalSemaphores[j] == 0x4000 + i * 100 + j;
		}
	}
	g_submit_ok = ok ? 1 : -1;
	return 0;
}

// Mirrors VkPhysicalDeviceMemoryProperties
typedef struct { uint32_t propertyFlags; uint32_t heapIndex; } MemoryType;
typedef struct { uint64_t size; uint32_t flags; } MemoryHeap;
typedef struct {
	uint32_t memoryTypeCount;
	MemoryType memoryTypes[32];
	uint32_t memoryHeapCount;
	MemoryHeap memoryHeaps[16];
} MemoryProperties;

void vkGetPhysicalDeviceMemoryProperties(uintptr_t pd, MemoryProperties *out) {
	usleep(g_delay_us);
	memset(out, 0, sizeof *out);
	out->memoryTypeCount = 3;
	out->memoryTypes[2].heapIndex = 7;
	out->memoryHeapCount = 2;
	out->memoryHeaps[1].size = 0x123456789ull;
}

// ---- support for resolve_test.go ----

// Instances get distinct handles so tests can tell them apart
static uintptr_t g_next_instance = 0xABC000;
int32_t vkCreateInstance(const void *info, const void *allocator, uintptr_t *out) { *out = g_next_instance++; return 0; }
void vkDestroyInstance(uintptr_t instance, const void *allocator) {}

// Not an exported symbol: reachable only through vkGetInstanceProcAddr, as an extension command is through the loader.
static void hiddenCmdSetStencilReference(uintptr_t cb, uint32_t mask, uint32_t ref) { g_i[0] = cb; g_i[1] = mask; g_i[2] = ref; }

// Like the loader: with no instance only global commands are returned (none here are, so everything falls through to
// dlsym); with an instance, commands behind the hidden function table are available.
void *vkGetInstanceProcAddr(uintptr_t instance, const char *name) {
	if (instance == 0) return NULL;
	if (strcmp(name, "vkCmdSetStencilReference") == 0) return (void *)hiddenCmdSetStencilReference;
	// The only name a device that has the extension, but not the core version, would offer for the same command
	if (strcmp(name, "vkCmdSetStencilReferenceTestEXT") == 0) return (void *)hiddenCmdSetStencilReference;
	return NULL;
}

uintptr_t testGetFloatBits(uintptr_t i) { uint32_t u; memcpy(&u, &g_f[i], sizeof u); return u; }
uintptr_t testGetInt(uintptr_t i) { return (uintptr_t)g_i[i]; }
void testReset(void) { memset(g_f, 0, sizeof g_f); memset(g_i, 0, sizeof g_i); }
`

var (
	cmdGetFloatBits = &vkCommand{"testGetFloatBits", 1, true, nil, nil}
	cmdGetInt       = &vkCommand{"testGetInt", 1, true, nil, nil}
	cmdSetDelay     = &vkCommand{"testSetDelay", 1, false, nil, nil}
	cmdSubmitOK     = &vkCommand{"testSubmitOK", 1, true, nil, nil}
	// testReset takes no arguments, but no generated command has an empty shape; the extra argument is ignored by the callee
	cmdReset = &vkCommand{"testReset", 1, false, nil, nil}
)

func gotFloat(i uintptr) float32 {
	return math.Float32frombits(uint32(vkCall_i(cmdGetFloatBits, i)))
}

func gotInt(i uintptr) uintptr { return vkCall_i(cmdGetInt, i) }

// reexecEnv guards against re-executing more than once
const reexecEnv = "VK_TEST_REEXEC"

// loaderEnv names the environment variable that opts in to the real-loader integration tests in loader_test.go
const loaderEnv = "VK_TEST_LOADER"

// skipReason is non-empty when the fake Vulkan library could not be built; tests that need it report SKIP with this reason.
var skipReason string

// requireFakeLibrary skips the calling test or benchmark if the fake Vulkan library is unavailable.
func requireFakeLibrary(tb testing.TB) {
	tb.Helper()
	if skipReason != "" {
		tb.Skip(skipReason)
	}
}

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		// The fake library is built with a Unix style C toolchain and its stubs use POSIX calls. There is no Windows
		// equivalent yet, so these tests are skipped (visibly, with -v) rather than run against a real vulkan-1.dll.
		skipReason = "the fake Vulkan library is not supported on windows"
		os.Exit(m.Run())
	}

	// The memory-safety tests in escapes_test.go need runtime debugging help to be reliable:
	//  - clobberfree=1 overwrites freed objects, so a use-after-free is seen as corruption rather than passing by luck
	//  - adaptivestackstart=0 gives every goroutine the same small initial stack, so stack growth is reproducible
	// GODEBUG is read at process start, so re-execute the test binary with it set.
	if os.Getenv(reexecEnv) == "" {
		child := exec.Command(os.Args[0], os.Args[1:]...)
		child.Stdin, child.Stdout, child.Stderr = os.Stdin, os.Stdout, os.Stderr
		godebug := strings.Trim(os.Getenv("GODEBUG")+",clobberfree=1,adaptivestackstart=0", ",")
		child.Env = append(os.Environ(), "GODEBUG="+godebug, reexecEnv+"=1")
		if err := child.Run(); err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				os.Exit(exitErr.ExitCode())
			}
			println("could not re-execute test binary:", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
	// Integration tests against a real Vulkan loader (see loader_test.go) use that library instead of the fake one.
	if loader := os.Getenv(loaderEnv); loader != "" {
		skipReason = "the fake Vulkan library is not used when " + loaderEnv + " is set"
		OverrideDefaultVulkanLibrary(loader)
		os.Exit(m.Run())
	}

	cc, err := exec.LookPath("cc")
	if err != nil {
		skipReason = "no C compiler (cc) found on PATH"
		os.Exit(m.Run())
	}

	dir, err := os.MkdirTemp("", "vkfake")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	src := filepath.Join(dir, "fake.c")
	if err := os.WriteFile(src, []byte(fakeVulkanC), 0666); err != nil {
		panic(err)
	}

	lib := filepath.Join(dir, "libfakevulkan.so")
	args := []string{"-shared", "-fPIC", "-o", lib, src}
	if runtime.GOOS == "darwin" {
		args = []string{"-dynamiclib", "-o", lib, src}
	}
	if out, err := exec.Command(cc, args...).CombinedOutput(); err != nil {
		println("could not build fake Vulkan library:", string(out))
		os.Exit(1)
	}

	OverrideDefaultVulkanLibrary(lib)
	os.Exit(m.Run())
}

func TestFloatArgs(t *testing.T) {
	requireFakeLibrary(t)

	tests := []struct {
		name   string
		call   func()
		ints   []uintptr
		floats []float32
	}{
		{"CmdSetLineWidth", func() { CmdSetLineWidth(CommandBuffer(0x1234), 2.5) },
			[]uintptr{0x1234}, []float32{2.5}},
		{"CmdSetDepthBias", func() { CmdSetDepthBias(CommandBuffer(0x1234), 1.5, -2.25, 3.75) },
			[]uintptr{0x1234}, []float32{1.5, -2.25, 3.75}},
		{"CmdSetDepthBounds", func() { CmdSetDepthBounds(CommandBuffer(0x1234), 0.25, 0.75) },
			[]uintptr{0x1234}, []float32{0.25, 0.75}},
		{"SetDeviceMemoryPriorityEXT", func() { SetDeviceMemoryPriorityEXT(Device(0x1234), DeviceMemory(0x5678), 0.5) },
			[]uintptr{0x1234, 0x5678}, []float32{0.5}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			vkCall_i(cmdReset, 0)
			tc.call()

			for i, want := range tc.ints {
				if got := gotInt(uintptr(i)); got != want {
					t.Errorf("integer arg %d: got %#x, want %#x", i, got, want)
				}
			}
			for i, want := range tc.floats {
				if got := gotFloat(uintptr(i)); got != want {
					t.Errorf("float arg %d: got %v, want %v", i, got, want)
				}
			}
		})
	}
}

// TestIntegerOnlyShapeLosesFloats is a negative control documenting why shapes exist: calling a float-taking function
// through an all-integer trampoline puts the float bits in an integer register, where the callee does not look for them.
func TestIntegerOnlyShapeLosesFloats(t *testing.T) {
	requireFakeLibrary(t)

	wrong := &vkCommand{"vkCmdSetLineWidth", 2, false, nil, nil}

	vkCall_i(cmdReset, 0)
	vkCall_ii(wrong, 0x1234, uintptr(math.Float32bits(2.5)))

	if got := gotFloat(0); got == 2.5 {
		t.Fatalf("float arrived intact through an integer-only trampoline; shapes may no longer be necessary on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
}
