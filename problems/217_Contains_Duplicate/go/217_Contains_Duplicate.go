package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func containsDuplicate(nums []int) bool {
	tmp := make(map[int]bool, 0)
	for _, num := range nums {
		_, found := tmp[num]
		if found {
			return true
		}

		tmp[num] = true
	}
	return false
}
