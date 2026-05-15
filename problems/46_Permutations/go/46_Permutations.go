package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func permute(nums []int) [][]int {
	var res [][]int
	var path []int

	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()

	return res
}
