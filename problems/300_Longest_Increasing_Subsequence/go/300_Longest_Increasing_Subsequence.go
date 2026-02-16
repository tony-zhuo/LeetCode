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

	/*
		1. 初始時 `dp = []`
		2. 處理 10：`dp = [10]`
		3. 處理 9：9 < 10，替換得 `dp = [9]`
		4. 處理 2：2 < 9，替換得 `dp = [2]`
		5. 處理 5：5 > 2，添加得 `dp = [2, 5]`
		6. 處理 3：3 > 2 但 3 < 5，替換得 `dp = [2, 3]`
		7. 處理 7：7 > 3，添加得 `dp = [2, 3, 7]`
		8. 處理 101：101 > 7，添加得 `dp = [2, 3, 7, 101]`
		9. 處理 18：18 < 101，替換得 `dp = [2, 3, 7, 18]`

	*/

	dp := make([]int, 0, len(nums))
	for _, num := range nums {
		dpIndex := sort.SearchInts(dp, num)
		if dpIndex == len(dp) {
			// 如果 `dpIndex == len(dp)`，說明 `num` 比 `dp` 中所有元素都大
			// 可以擴展目前的最長遞增子序列，將 `num` 添加到 `dp` 末尾
			dp = append(dp, num)
		} else {
			dp[dpIndex] = num
		}
	}

	return len(dp)
}
