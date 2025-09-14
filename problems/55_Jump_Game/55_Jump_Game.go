package problems

func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])

		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return maxReach >= len(nums)-1
}
