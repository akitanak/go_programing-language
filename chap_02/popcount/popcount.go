package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popcountはxのポピュレーションカウント(1が設定されているビット数)を返します。
func PopCount(x uint64) int {
	// オリジナル実装
	// return int(pc[byte(x>>(0*8))] +
	// 	pc[byte(x>>(1*8))] +
	// 	pc[byte(x>>(2*8))] +
	// 	pc[byte(x>>(3*8))] +
	// 	pc[byte(x>>(4*8))] +
	// 	pc[byte(x>>(5*8))] +
	// 	pc[byte(x>>(6*8))] +
	// 	pc[byte(x>>(7*8))])
	const maxByteCount = 8
	var count int
	for i := 0; i < maxByteCount; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func InefficientPopCount(x uint64) int {
	count := x & 1
	for i := 0; i < 64; i++ {
		x = x >> 1
		count += x & 1
	}
	return int(count)
}

func AnotherPopCount(x uint64) int {
	var count int
	for x > 0 {
		x = x & (x - 1)
		count++
	}

	return count
}
