// スライスの代わりに配列へのポインタを使うようにreverseを書き直す。
package main

import (
	"fmt"
)

func main() {
	slice := [5]int{1, 2, 3, 4, 5}
	reverse(&slice)
	fmt.Println(slice)
}

func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
