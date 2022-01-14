// 入力テキストファイル内のそれぞれの単語の出現頻度を報告するプログラムwordfreqを書きなさい。入力を行ではなく単語へ分割するために、最初のScan呼び出しの前にinput.Split(bufio.ScanWords)を呼び出しなさい。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		occurrences := wordfreq(os.Stdin)
		printOccurrences(occurrences)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
				continue
			}
			defer f.Close()

			occurrences := wordfreq(f)
			printOccurrences(occurrences)
		}
	}
}

func wordfreq(f *os.File) map[string]int {
	occurrences := make(map[string]int)

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		occurrences[input.Text()]++
	}

	return occurrences
}

func printOccurrences(occurrences map[string]int) {
	fmt.Println("word\tcount")
	for w, count := range occurrences {
		fmt.Printf("%v\t%d\n", w, count)
	}
}
