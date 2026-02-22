package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap
func scoreDifference(nums []int) int {
	firstPlayer := 0
	secondPlayer := 0
	current := 1
	for game, num := range nums {
		if num%2 != 0 {
			current = 3 - current
		}
		if (game+1)%6 == 0 {
			current = 3 - current
		}
		if current == 1 {
			firstPlayer += num
		} else {
			secondPlayer += num
		}
	}

	return firstPlayer - secondPlayer
}
