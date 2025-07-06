package problems

var tmp map[int]int = map[int]int{}

func ClimbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if num, ok := tmp[n]; ok {
		return num
	}

	tmp[n] = ClimbStairs(n-1) + ClimbStairs(n-2)
	return tmp[n]
}
