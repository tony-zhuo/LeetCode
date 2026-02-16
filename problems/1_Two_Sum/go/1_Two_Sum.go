package problems

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))

	for current_index, current_num := range nums {
		if target_index, found := numMap[target-current_num]; found {
			return []int{target_index, current_index}
		}

		if _, found := numMap[current_num]; !found {
			numMap[current_num] = current_index
		}
	}

	return []int{}
}
