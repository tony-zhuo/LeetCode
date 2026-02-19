package problems

func climbStairs_hashmap(n int) int {
	dp := make(map[int]int, n)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs_dp(n int) int {
	a, b := 1, 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

type Matrix [2][2]int

func multiply(a, b Matrix) Matrix {
	return Matrix{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
	}
}
func matpow(m Matrix, n int) Matrix {
	// 單位矩陣（相當於數字的 1）
	result := Matrix{{1, 0}, {0, 1}}
	for n > 0 {
		if n%2 == 1 {
			result = multiply(result, m)
		}
		m = multiply(m, m) // 平方
		n /= 2
	}
	return result
}
func climbStairs_matrix(n int) int {
	if n <= 1 {
		return 1
	}
	base := Matrix{{1, 1}, {1, 0}}
	res := matpow(base, n)
	return res[0][0]
}
