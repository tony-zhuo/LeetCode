package problems

func reverseBits(n int) int {
	var result int
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= (n & 1)
		n >>= 1
	}
	return result
}
