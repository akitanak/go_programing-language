// 一回のパスで操作を行うrotateを書く
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = rotate(s, 3)
	fmt.Println(s)
}

func rotate(s []int, step int) []int {
	front := s[:step]
	rest := s[step:]
	return append(rest, front...)
}
