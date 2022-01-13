// UTF-8でエンコードされた文字列を表す[]byteスライスの文字を、そのスライス内で逆順にするようにreverseを修正しなさい。新たなメモリを割り当てることなく行えるでしょうか？
package main

import (
	"fmt"
	"os"
)

func main() {
	s := []byte("abcあいうdefえおか")
	actual := reverse(&s)
	want := []byte("かおえfedういあcba")
	if string(actual) != string(want) {
		fmt.Printf("expected: %X, actual: %X\n", want, actual)
		os.Exit(1)
	}
	fmt.Printf("result: %v\n", string(actual))
}

func reverse(s *[]byte) []byte {
	origin := *s
	result := make([]byte, len(*s))
	resultIndex := len(result)
	for i, b := range origin {
		byteCount := detectBytesCount(b)
		if byteCount == 0 {
			continue
		}
		char := origin[i : i+byteCount]
		resultIndex -= len(char)
		copyChar(&result, resultIndex, char)
	}

	return result
}

func detectBytesCount(b byte) int {
	if b >= 0x00 && b <= 0x7f {
		return 1
	} else if b >= 0xc2 && b <= 0xdf {
		return 2
	} else if b >= 0xe0 && b <= 0xef {
		return 3
	} else if b >= 0xf0 && b <= 0xf4 {
		return 4
	}
	return 0
}

func copyChar(to *[]byte, index int, char []byte) {
	for i := 0; i < len(char); i++ {
		(*to)[index+i] = char[i]
	}
}
