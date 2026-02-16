package problems

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	sum := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = 1
		} else {
			tmp := nums[i-1] * sum
			res[i] = tmp
			sum = tmp
		}
	}

	sum = 1
	for i := len(nums) - 2; i >= 0; i-- {
		tmp := nums[i+1] * sum
		res[i] *= tmp
		sum = tmp
	}
	return res
}

/*
tmp 1
(1) 1  2   3  4
	1  1   2  6 (1)
    24 12  8  6
*/
