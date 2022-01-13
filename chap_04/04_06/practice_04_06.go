// UTF-8でエンコードされた[]byteスライス内で隣接しているUnicodeスペース(unicode.IsSpaceを参照)をもとのスライス内で一つのASCIIスペースへ圧縮する関数を書く。
package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	s := []byte("ab \tcdefg\nhijk\v\f \u00A0lmnop\r\ns\u0085\t")
	fmt.Printf("original: %s\n", s)
	s = compressSpace(s)
	fmt.Printf("result: %s\n", s)
	expected := "ab cdefg hijk lmnop s "
	if string(s) != expected {
		fmt.Printf("expected: %X, actual: %X\n", expected, s)
		os.Exit(1)
	}
}

func compressSpace(s []byte) []byte {
	var existStartSpace bool
	var i int
	// ひとまずunicodeスペースをASCIIスペースへ変換する。
	for _, r := range s {
		if unicode.IsSpace(rune(r)) {
			// スペースの直前がC2だった場合は、C2を除去するためにインデックスを一つ戻す。
			if i > 0 && s[i-1] == 0xc2 {
				i--
			}
			if existStartSpace {
				continue
			} else {
				s[i] = 32
				existStartSpace = true
			}
		} else {
			s[i] = r
			if rune(r) == 0xc2 {
				continue
			}
			existStartSpace = false
		}
		i++
	}
	return s[:i]
}
