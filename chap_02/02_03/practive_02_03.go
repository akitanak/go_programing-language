package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"local.packages/chap_02/popcount"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			v, err := strconv.ParseUint(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "practice 2.2: %v\n", err)
				os.Exit(1)
			}
			count := popcount.PopCount(v)
			fmt.Printf("%v contains %d\n", v, count)
		}
	} else {
		for _, arg := range args {
			v, err := strconv.ParseUint(arg, 10, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "practice 2.2: %v\n", err)
				os.Exit(1)
			}
			count := popcount.PopCount(v)
			fmt.Printf("%v contains %d\n", v, count)
		}
	}
}
