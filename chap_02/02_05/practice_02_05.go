package main

import (
	"fmt"
	"time"

	"local.packages/chap_02/popcount"
)

func main() {
	const v = 2 ^ 64
	const maxLoopCount = 300000
	var result int
	fmt.Println("---- PopCount ----")
	start := time.Now()
	for i := 0; i < maxLoopCount; i++ {
		result = popcount.PopCount(v)
	}
	fmt.Printf("answer: %d, durarion: %.5f\n", result, time.Since(start).Seconds())

	fmt.Println("---- InEfficientPopCount ----")
	start = time.Now()
	for i := 0; i < maxLoopCount; i++ {
		result = popcount.InefficientPopCount(v)
	}
	fmt.Printf("answer: %d, durarion: %.5f\n", result, time.Since(start).Seconds())

	fmt.Println("---- AnotherPopCount ----")
	start = time.Now()
	for i := 0; i < maxLoopCount; i++ {
		result = popcount.AnotherPopCount(v)
	}
	fmt.Printf("answer: %d, durarion: %.5f\n", result, time.Since(start).Seconds())
}
