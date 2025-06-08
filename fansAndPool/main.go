package main

import (
	"context"
	"fanin-fanout/pkg"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("FANIN-FANOUT")
	UseFanInFanOut(20)
	// fmt.Println("POOL")
	// UsePool(20)

}

func UseFanInFanOut(chCount int) {
	start := time.Now()
	in := make(chan int)

	m := runtime.MemStats{}

	go func() {
		defer close(in)
		for i := range 10000 {
			in <- i
		}
	}()

	ctx := context.Background()

	out := pkg.FanIn(ctx, pkg.FanOut(ctx, in, chCount), pkg.Procces1)

	runtime.ReadMemStats(&m)
	fmt.Println("Before Heap inuse: ", m.HeapInuse/1024)
	fmt.Println("Before Head alloc: ", m.HeapAlloc/1024)

	fmt.Println("Before goroutines: ", runtime.NumGoroutine())

	for range out {
	}

	runtime.GC()

	runtime.ReadMemStats(&m)
	fmt.Println("Final Heap inuse: ", m.HeapInuse/1024)
	fmt.Println("Final Head alloc: ", m.HeapAlloc/1024)
	fmt.Println("Final goroutines: ", runtime.NumGoroutine())

	fmt.Println("Sinced: ", time.Since(start))

}

func UsePool(wCount int) {
	start := time.Now()
	in := make(chan int)

	m := runtime.MemStats{}

	go func() {
		defer close(in)
		for i := range 10000 {
			in <- i
		}
	}()

	out := pkg.PoolWokers(in, wCount, pkg.Procces1)

	runtime.ReadMemStats(&m)
	fmt.Println("Before Heap inuse: ", m.HeapInuse/1024)
	fmt.Println("Before Head alloc: ", m.HeapAlloc/1024)

	for range out {
	}

	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Println("Final Heap inuse: ", m.HeapInuse/1024)
	fmt.Println("Final Head alloc: ", m.HeapAlloc/1024)

	fmt.Println("Sinced: ", time.Since(start))

}
