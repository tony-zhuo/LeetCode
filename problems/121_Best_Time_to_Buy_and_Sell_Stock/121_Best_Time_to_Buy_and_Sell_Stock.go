package problems

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	left := 0
	bestProfit := 0

	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			left = right
		} else {
			bestProfit = max(bestProfit, prices[right]-prices[left])
		}
	}

	if bestProfit < 0 {
		return 0
	}
	return bestProfit
}
