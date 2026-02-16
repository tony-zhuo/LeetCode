package problems

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		numMap[num] = true
	}

	res := 0

	for num, _ := range numMap {
		if numMap[num+1] {
			continue
		}

		tmp := 0
		for cur := num; numMap[cur]; cur-- {
			tmp++
		}

		res = max(tmp, res)
	}

	return res
}
