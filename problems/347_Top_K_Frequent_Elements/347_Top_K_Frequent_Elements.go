package problems

import "sort"

func topKFrequent(nums []int, k int) []int {
	tmpSet := make(map[int]int, len(nums))
	for _, num := range nums {
		tmpSet[num]++
	}

	tmpArr := make([][]int, 0, len(tmpSet))
	for num, count := range tmpSet {
		tmpArr = append(tmpArr, []int{num, count})
	}

	sort.Slice(tmpArr, func(i, j int) bool {
		return tmpArr[i][1] > tmpArr[j][1]
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, tmpArr[i][0])
	}

	return res
}
