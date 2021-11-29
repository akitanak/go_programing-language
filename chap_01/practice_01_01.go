package main

import (
	"fmt"
	"os"
)

// echoプログラムを修正してそのプログラムを起動したコマンド名であるos.Args[0]も表示するようにする。
func main() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
