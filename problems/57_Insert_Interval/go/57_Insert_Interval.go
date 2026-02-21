package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			// interval 完全在 newInterval 左邊
			res = append(res, interval)
		} else if interval[0] > newInterval[1] {
			// interval 完全在 newInterval 右邊
			res = append(res, newInterval)
			newInterval = interval
		} else {
			// interval 與 newInterval 重合
			if interval[0] < newInterval[0] {
				newInterval[0] = interval[0]
			}
			if interval[1] > newInterval[1] {
				newInterval[1] = interval[1]
			}
		}
	}

	res = append(res, newInterval)

	return res
}
