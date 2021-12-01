package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	length := len(s)
	if length <= 3 {
		return s
	}

	var buffer bytes.Buffer
	rest := length % 3
	if rest > 0 {
		fmt.Fprintf(&buffer, "%s,", s[:rest])
	}
	for i := rest; i < length; i += 3 {
		fmt.Fprintf(&buffer, "%s", s[i:i+3])
		if i+3 < length {
			fmt.Fprint(&buffer, ",")
		}
	}
	return buffer.String()
}
