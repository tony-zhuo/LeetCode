package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	globalMax := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		globalMax = max(currentSum, globalMax)
	}

	return globalMax
}
