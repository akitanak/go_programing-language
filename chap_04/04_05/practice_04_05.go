// []stringスライス内で隣接している重複をスライス内で除去する関数を書く。
package main

import "fmt"

func main() {
	s := []string{"one", "two", "three", "three", "four", "five", "six", "six", "seven", "eight", "eight", "nine"}

	s = removeDup(s)
	fmt.Println(s)
}

func removeDup(s []string) []string {
	var prev string
	var dups []int
	for i, v := range s {
		if v == prev {
			dups = append(dups, i)
		}
		prev = v
	}
	for i, dup := range dups {
		s = remove(s, dup-i)
	}
	return s
}

func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
