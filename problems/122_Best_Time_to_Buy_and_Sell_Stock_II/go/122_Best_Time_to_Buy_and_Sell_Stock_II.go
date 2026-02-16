package problems

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	left := 0
	profit := 0

	for right := 1; right < len(prices); right++ {
		if prices[right] > prices[left] {
			profit += prices[right] - prices[left]
			left = right
		} else {
			left = right
		}
	}

	return profit
}
