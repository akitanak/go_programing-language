package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 非効率な可能性のあるバージョンとstrings.Joinを使ったバージョンとで、実行時間の差を計測しなさい。
func inefficiencyEcho() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func efficiencyEcho() {
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}

func monitorPerformance(fn func()) time.Duration {
	start := time.Now()
	for i := 0; i < 100; i++ {
		fn()
	}
	return time.Since(start)
}

func main() {
	duration1 := monitorPerformance(inefficiencyEcho)
	fmt.Printf("inefficiencyEcho: %d\n", duration1)

	duration2 := monitorPerformance(efficiencyEcho)
	fmt.Printf("efficiencyEcho: %d\n", duration2)

	fmt.Printf("inefficiencyEcho: %d, efficiencyEcho: %d, diff: %d\n", duration1, duration2, duration1-duration2)
}
