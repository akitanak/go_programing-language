package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

func main() {
	method := "sha256"
	if len(os.Args) >= 2 {
		method = strings.ToLower(os.Args[1])
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		b := scanner.Bytes()
		switch method {
		case "sha256":
			hash := sha256.Sum256(b)
			fmt.Printf("%x\n", hash)
		case "sha384":
			hash := sha512.Sum384(b)
			fmt.Printf("%x\n", hash)
		case "sha512":
			hash := sha512.Sum512(b)
			fmt.Printf("%x\n", hash)
		default:
			hash := sha256.Sum256(b)
			fmt.Printf("%x\n", hash)
		}
	}

}
