package main

import (
	"fmt"
	"os"
)

// echoプログラムを修正して,個々の引数のインデックスと値の組を1行ごとに表示しなさい。
func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
