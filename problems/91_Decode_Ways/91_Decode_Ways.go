package problems

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	//fmt.Println(int('0'))
	// '0' = 48
	// '1' = 49
	// '2' = 50

	dp := make([]int, n+1)
	dp[0] = 1 // 虛擬字

	// 第一個字
	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}

	for i := 2; i <= n; i++ {
		// 單一個字，僅檢查是否等於 0
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		two := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if two >= 10 && two <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
