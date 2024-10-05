package problems

import "sort"

func groupAnagrams(strs []string) [][]string {
	tmp := make(map[string][]string, len(strs))
	res := make([][]string, 0, len(strs))

	for _, str := range strs {
		sortStr := []rune(str)
		sort.Slice(sortStr, func(i, j int) bool {
			return sortStr[i] < sortStr[j]
		})

		tmp[string(sortStr)] = append(tmp[string(sortStr)], str)
	}

	for _, item := range tmp {
		res = append(res, item)
	}

	return res
}
