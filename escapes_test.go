package vk

import (
	"runtime"
	"sync"
	"testing"
	"unsafe"
)

// Commands pass Go pointers to C as uintptr(unsafe.Pointer(x)). The runtime does not track a uintptr, which creates two
// hazards that the //go:uintptrescapes directive on every vkCall_* wrapper exists to remove:
//
//  1. Liveness: after the conversion, nothing references the target, so the GC may free (and reuse) it while a long-running
//     command is still reading it.
//  2. Stack movement: a target on the goroutine stack is not updated when the stack is copied to grow it, which can happen
//     in the prologue of any function called after the uintptr is made.
//
// Both tests below are deterministic reproductions; they fail against wrappers without the directive.

func gotSubmitOK() int { return int(int32(vkCall_i(cmdSubmitOK, 0))) }

func setDelayMicros(us uintptr) { vkCall_i(cmdSetDelay, us) }

// makeSubmits builds a pointer graph (slices of handles, plus a pNext target) that is referenced only by the returned
// slice. The pattern must match vkQueueSubmit in the fake library.
//
//go:noinline
func makeSubmits(count int) []SubmitInfo {
	submits := make([]SubmitInfo, count)
	for i := range submits {
		s := &submits[i]
		s.PWaitSemaphores = make([]Semaphore, 64)
		s.PWaitDstStageMask = make([]PipelineStageFlags, 64)
		s.PCommandBuffers = make([]CommandBuffer, 64)
		s.PSignalSemaphores = make([]Semaphore, 64)
		for j := 0; j < 64; j++ {
			s.PWaitSemaphores[j] = Semaphore(0x1000 + i*100 + j)
			s.PWaitDstStageMask[j] = PipelineStageFlags(0x2000 + i*100 + j)
			s.PCommandBuffers[j] = CommandBuffer(0x3000 + i*100 + j)
			s.PSignalSemaphores[j] = Semaphore(0x4000 + i*100 + j)
		}
		next := new(uint64)
		*next = 0xFEEDFACE00000000 + uint64(i)
		s.PNext = unsafe.Pointer(next)
	}
	return submits
}

// churn forces garbage collections and allocates over freed memory until stopped.
func churn(stop <-chan struct{}) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var keep [][]uint64
		for {
			select {
			case <-stop:
				return
			default:
			}
			runtime.GC()
			keep = keep[:0]
			for i := 0; i < 2000; i++ {
				b := make([]uint64, 64)
				for j := range b {
					b[j] = 0xCDCDCDCDCDCDCDCD
				}
				keep = append(keep, b)
			}
		}
	}()
	return wg
}

func TestPointerArgsSurviveGCDuringCall(t *testing.T) {
	requireFakeLibrary(t)

	if os := runtime.GOMAXPROCS(4); os < 4 {
		defer runtime.GOMAXPROCS(os)
	}

	setDelayMicros(20_000)
	defer setDelayMicros(0)

	stop := make(chan struct{})
	wg := churn(stop)
	defer func() { close(stop); wg.Wait() }()

	for i := 0; i < 25; i++ {
		// The slice is built in a function and passed directly, so the caller holds no reference to it (or anything it
		// points to) during the call; the only references are the pointers handed to C.
		if r := QueueSubmit(Queue(0x51), makeSubmits(4), Fence(0xFE)); r != SUCCESS {
			t.Fatalf("QueueSubmit returned %v", r)
		}
		if ok := gotSubmitOK(); ok != 1 {
			t.Fatalf("iteration %d: the command saw corrupted or freed input (verdict %d)", i, ok)
		}
	}
}

//go:noinline
func recurse(depth int, f func()) byte {
	var pad [56]byte
	pad[depth&31] = byte(depth)
	if depth > 0 {
		pad[0] += recurse(depth-1, f)
	} else {
		f()
	}
	return pad[0] + pad[depth&31]
}

//go:noinline
func recurseWide(depth int, f func()) byte {
	var pad [200]byte
	pad[depth&31] = byte(depth)
	if depth > 0 {
		pad[0] += recurseWide(depth-1, f)
	} else {
		f()
	}
	return pad[0] + pad[depth&31]
}

func TestOutputPointerSurvivesStackGrowth(t *testing.T) {
	requireFakeLibrary(t)

	check := func() bool {
		mp := GetPhysicalDeviceMemoryProperties(PhysicalDevice(1))
		return mp.MemoryTypeCount == 3 && mp.MemoryTypes[2].HeapIndex == 7 &&
			mp.MemoryHeapCount == 2 && mp.MemoryHeaps[1].Size == 0x123456789
	}

	// Call the command with the goroutine's stack consumed to every depth in a range, each on a fresh goroutine with a
	// small, fixed initial stack. Somewhere in the sweep the stack runs out between the command's own prologue and the
	// prologues of the calls it makes to reach C; if the output buffer is on the stack, C writes to the abandoned old copy.
	var failures []int
	for _, rec := range []func(int, func()) byte{recurse, recurseWide} {
		for depth := 0; depth < 400; depth++ {
			ok := make(chan bool)
			go func() { rec(depth, func() { ok <- check() }) }()
			if !<-ok {
				failures = append(failures, depth)
			}
		}
	}

	if len(failures) > 0 {
		first := failures
		if len(first) > 5 {
			first = first[:5]
		}
		t.Errorf("output of GetPhysicalDeviceMemoryProperties was lost at %d stack depths (first few: %v)", len(failures), first)
	}
}

func BenchmarkCmdSetLineWidth(b *testing.B) {
	requireFakeLibrary(b)
	for i := 0; i < b.N; i++ {
		CmdSetLineWidth(CommandBuffer(1), 2.5)
	}
}

func BenchmarkGetPhysicalDeviceMemoryProperties(b *testing.B) {
	requireFakeLibrary(b)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetPhysicalDeviceMemoryProperties(PhysicalDevice(1))
	}
}

func BenchmarkQueueSubmit(b *testing.B) {
	requireFakeLibrary(b)
	b.ReportAllocs()
	submits := makeSubmits(1)
	for i := 0; i < b.N; i++ {
		QueueSubmit(Queue(0x51), submits, Fence(0xFE))
	}
}
