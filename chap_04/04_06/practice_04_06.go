// UTF-8でエンコードされた[]byteスライス内で隣接しているUnicodeスペース(unicode.IsSpaceを参照)をもとのスライス内で一つのASCIIスペースへ圧縮する関数を書く。
package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []byte("ab \tcdefg\nhijk\v\f lmnop\r\ns\u0085\t")
	fmt.Printf("original: %s\n", s)
	s = compressSpace(s)
	fmt.Printf("result: %s\n", s)
}

func compressSpace(s []byte) []byte {
	var existStartSpace bool
	var i int
	// ひとまずunicodeスペースをASCIIスペースへ変換する。
	for _, r := range s {
		if unicode.IsSpace(rune(r)) {
			if existStartSpace {
				continue
			} else {
				// 2byte文字のスペース(u+0085, U+00A0)だった場合は、直前のbyteのC2を除去するためにindexを戻す。
				if i > 0 && s[i-1] == 194 {
					i--
				}
				s[i] = 32
				existStartSpace = true
			}
		} else {
			s[i] = r
			existStartSpace = false
		}
		i++
	}
	return s[:i]
}
