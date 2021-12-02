package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(detectAnagram("anagrams", "ARS MAGNA"))
	fmt.Println(detectAnagram("Statue of Liberty", "built to stay free"))
	fmt.Println(detectAnagram("Christmas", "trims cash"))
	fmt.Println(detectAnagram("astronomer", "moon starer"))
	fmt.Println(detectAnagram("天地創造", "天創　地造"))
	fmt.Println(detectAnagram("たけやぶ", "やぶけた"))
	fmt.Println(detectAnagram("anagrams", "slam gana"))
}

func detectAnagram(s1 string, s2 string) bool {
	s1, s2 = removeBlank(s1), removeBlank(s2)
	s1, s2 = strings.ToLower(s1), strings.ToLower(s2)
	s1, s2 = sortString(s1), sortString(s2)
	return s1 == s2
}

func removeBlank(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if r != ' ' && r != '　' {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func sortString(s string) string {
	var ints []int
	for _, r := range s {
		ints = append(ints, int(r))
	}

	sort.Ints(ints)

	var builder strings.Builder
	for _, v := range ints {
		builder.WriteRune(rune(v))
	}
	return builder.String()
}
