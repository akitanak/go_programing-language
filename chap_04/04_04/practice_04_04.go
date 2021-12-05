// 一回のパスで操作を行うrotateを書く
package main

func main() {

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i<j; i, j + i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, step int) {
	reverse(s[:step])
	reverse(s[step:])
	reverse(s)
}
