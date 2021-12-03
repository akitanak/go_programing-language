package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	x := sha256.Sum256([]byte("X"))
	y := sha256.Sum256([]byte("Y"))

	fmt.Println(countDifferentBit(x, y))
}

// 2つのSHA256ハッシュで異なるビットの数を数える
func countDifferentBit(x, y [32]byte) int {
	var count int
	for i := 0; i < len(x); i++ {
		diff := uint8(x[i]) ^ uint8(y[i])
		count += popcount(diff)
	}
	return count
}

func popcount(x uint8) int {
	var count int
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
