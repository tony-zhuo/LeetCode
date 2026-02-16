package problems

//func missingNumber(nums []int) int {
//	n := len(nums)
//
//	src := 0
//	for i := 0; i <= n; i++ {
//		src += i
//	}
//
//	numSum := 0
//	for _, num := range nums {
//		numSum += num
//	}
//
//	return src - numSum
//}

func missingNumber(nums []int) int {
	result := len(nums)
	for i, num := range nums {
		result = result ^ i ^ num
	}

	return result
}
