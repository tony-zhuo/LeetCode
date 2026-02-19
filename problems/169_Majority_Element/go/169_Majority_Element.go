package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func majorityElement(nums []int) int {
	candidate := 0
	cnt := 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if candidate == num {
			cnt += 1
		} else {
			cnt -= 1
		}
	}

	return candidate
}
