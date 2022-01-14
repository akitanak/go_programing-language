// unicode.IsLetterなどの関数を使ってUnicode分類に従って文字や数字などを数えるようにcharcountを修正する。
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type class string

const (
	Control = class("control")
	Letter  = class("letter")
	Mark    = class("mark")
	Number  = class("number")
	Punct   = class("puct")
	Space   = class("space")
	Symbol  = class("symbol")
)

func main() {
	classCounts := make(map[class]int)
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		c, err := classificateRune(r)
		if err == nil {
			classCounts[c]++
		} else {
			fmt.Printf("charcount: %v\n", err)
		}
	}
	fmt.Print("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nlen\tcount\n")
	for c, n := range classCounts {
		fmt.Printf("%v\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func classificateRune(r rune) (class, error) {
	switch {
	case unicode.IsControl(r):
		return Control, nil
	case unicode.IsLetter(r):
		return Letter, nil
	case unicode.IsMark(r):
		return Mark, nil
	case unicode.IsNumber(r):
		return Number, nil
	case unicode.IsPunct(r):
		return Punct, nil
	case unicode.IsSpace(r):
		return Space, nil
	case unicode.IsSymbol(r):
		return Symbol, nil
	default:
		return "", fmt.Errorf("%X is undefined class", r)
	}
}
