package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("test 1 Hello, world!")
	// testCPUBenchmark()
}

func testCPUBenchmark() {
	start := time.Now()
	var x float64 = 1.0001
	for i := 0; i < 1e8; i++ {
		x = x * 1.000001
	}
	elapsed := time.Since(start)
	fmt.Printf("CPU benchmark: %.2fs\n", elapsed.Seconds())
}

func testMemoryBenchmark() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	startAlloc := mem.Alloc

	slice := make([]byte, 100*1024*1024) // 100MB
	for i := range slice {
		slice[i] = byte(i)
	}

	runtime.ReadMemStats(&mem)
	endAlloc := mem.Alloc
	fmt.Printf("Memory allocated: %.2f MB\n", float64(endAlloc-startAlloc)/1024/1024)
}

func init() {
	testHeavyCPUBenchmark()
	testCPUBenchmark()
	testMemoryBenchmark()

}

func testHeavyCPUBenchmark() {
	start := time.Now()
	var result float64 = 1.0
	for j := 0; j < 8; j++ { // Run multiple times to increase load
		x := 1.000001
		for i := 0; i < 5e8; i++ { // Increase iterations for heavier load
			x = x * 1.000001
			if x > 1e10 {
				x = 1.000001
			}
		}
		result += x
	}
	elapsed := time.Since(start)
	fmt.Printf("Heavy CPU benchmark: %.2fs (result=%.2f)\n", elapsed.Seconds(), result)
}
