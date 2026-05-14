package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap
import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var path []int

	var backtrack func(start, remaining int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				break
			}

			path = append(path, candidates[i])
			backtrack(i, remaining-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)

	return res
}
