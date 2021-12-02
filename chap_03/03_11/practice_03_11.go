package main

import (
	"bytes"
	"fmt"
	"strings"
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
	fmt.Println(comma("-1234567890"))
	fmt.Println(comma("+1234567890"))
	fmt.Println(comma("1.234567890"))
	fmt.Println(comma("12.34567890"))
	fmt.Println(comma("123.4567890"))
	fmt.Println(comma("1234.567890"))
	fmt.Println(comma("12345.67890"))
	fmt.Println(comma("123456.7890"))
	fmt.Println(comma("-123456.7890"))
}

func comma(s string) string {
	var buffer bytes.Buffer
	// 符号の処理
	sign := string(s[0])
	if sign == "-" || sign == "+" {
		fmt.Fprint(&buffer, sign)
		s = s[1:]
	}

	// 小数点の処理
	var afterDecimalPointStr string
	decimalPointIndex := strings.IndexByte(s, '.')
	if decimalPointIndex != -1 {
		afterDecimalPointStr = s[decimalPointIndex:]
		s = s[:decimalPointIndex]
	}

	length := len(s)
	if length <= 3 {
		fmt.Fprint(&buffer, s)
		if afterDecimalPointStr != "" {
			fmt.Fprint(&buffer, afterDecimalPointStr)
		}
		return buffer.String()
	}

	rest := length % 3
	if rest > 0 {
		fmt.Fprintf(&buffer, "%s,", s[:rest])
	}
	for i := rest; i < length; i += 3 {
		fmt.Fprint(&buffer, s[i:i+3])
		if i+3 < length {
			fmt.Fprint(&buffer, ",")
		}
	}

	if afterDecimalPointStr != "" {
		fmt.Fprint(&buffer, afterDecimalPointStr)
	}
	return buffer.String()
}
