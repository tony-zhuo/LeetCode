package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		// nums[i] + left, right
		// left, right 要如何移動？
		// i 從 0 idx 開始
		// left = i+1
		// right = i + 1
		// 如果相加 > 0, right --
		// 如果相加 < 0, left ++
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right && left < len(nums) && right > 0 {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}
		}
	}

	return res
}
