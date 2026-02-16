package problems

func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatestNum := 0
	res := make([]bool, len(candies))

	for _, candy := range candies {
		greatestNum = max(candy, greatestNum)
	}

	for i, candy := range candies {
		if (candy + extraCandies) >= greatestNum {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res
}
