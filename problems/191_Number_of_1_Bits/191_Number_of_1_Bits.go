package problems

func hammingWeight(n int) int {
	cnt := 0
	// solution 1:

	//for i := 0; i <= 32; i++ {
	//	if (n & 1) == 1 {
	//		cnt++
	//	}
	//	n = n >> 1
	//}
	//
	//return cnt

	// solution 2:

	for n > 0 {
		cnt++
		n = n & (n - 1)
	}

	return cnt
}
