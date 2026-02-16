package problems

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[i]: 湊成金額 i 所需的 最少硬幣數
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
