package problems

import "sort"

func lengthOfLIS(nums []int) int {
	// solution 1：classic dp
	//n := len(nums)
	//dp := make([]int, len(nums))
	//ans := 0
	//
	//for i := 0; i < n; i++ {
	//	dp[i] = 1
	//	for j := 0; j < i; j++ {
	//		if nums[j] < nums[i] && dp[j]+1 > dp[i] {
	//			dp[i] = dp[j] + 1
	//		}
	//	}
	//	if dp[i] > ans {
	//		ans = dp[i]
	//	}
	//}
	//
	//return ans

	/*
		index ｜ nums[index] ｜ progress
		0     ｜ 10          ｜ 第一個，dp[0] = 1
		1     ｜ 9           ｜ 9 沒有比 10 大，dp[1] = 1
		2     ｜ 2           ｜ 前兩個都 > 2，dp[2] = 1
		3     ｜ 5           ｜ 只有 5 > 2，dp[3] = dp[2] + 1
		4     ｜ 3           ｜ 3 只有 > dp[2] 的 2，do[4] = d[2] + 1 = 2
		5     ｜ 7           ｜ 7 > 3, 5, 2，在 dp[5] = max(dp[2], dp[3], dp[4]) + 1 = 3
		6     ｜ 101         ｜ 101 大於前面全部，所以 dp[6] = dp[5](前面最大的)+1 = 4
		7     ｜ 18          ｜ 19 除了 101 都比較大，dp[7] = 最大的 dp[5] + 1 = 4
	*/

	// solution 2：Patience Sorting + 二分搜尋（O(n log n)）
	dp := make([]int, 0, len(nums))
	for _, num := range nums {
		dpIndex := sort.SearchInts(dp, num)
		if dpIndex == len(dp) {
			dp = append(dp, num)
		} else {
			dp[dpIndex] = num
		}
	}

	return len(dp)
}
