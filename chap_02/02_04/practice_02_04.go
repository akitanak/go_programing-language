package main

import (
	"fmt"
	"time"

	"local.packages/chap_02/popcount"
)

func main() {
	const v = 2 ^ 64
	const maxLoopCount = 300000

	fmt.Println("---- PopCount ----")
	start := time.Now()
	for i := 0; i < maxLoopCount; i++ {
		_ = popcount.PopCount(v)
	}
	fmt.Printf("durarion: %.5f\n", time.Since(start).Seconds())
	fmt.Println("---- InEfficientPopCount ----")
	start = time.Now()
	for i := 0; i < maxLoopCount; i++ {
		_ = popcount.InefficientPopCount(v)
	}
	fmt.Printf("durarion: %.5f\n", time.Since(start).Seconds())
}
