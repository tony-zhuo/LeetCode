package problems

func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 1; i <= n; i++ {
		cnt := 0
		tmp := i
		for tmp > 0 {
			cnt++
			tmp = tmp & (tmp - 1)
		}

		res[i] = cnt
	}

	return res
}
