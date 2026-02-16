package problems

/*
can not use extra space
use two pointer solve this question
already sorted
1-indexed
*/

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			right--
		}

		if sum < target {
			left++
		}
	}

	return []int{-1, -1}
}
