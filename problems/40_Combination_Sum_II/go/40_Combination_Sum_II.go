package problems

import "sort"

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	var path []int

	used := make([]bool, len(candidates))
	var backtrack func(start, num int)
	backtrack = func(start, num int) {
		if num == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if used[i] {
				return
			}
			if candidates[i] > num {
				return
			}

			used[i] = true
			path = append(path, candidates[i])

			backtrack(i+1, num-candidates[i])

			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)

	return res
}
