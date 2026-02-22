package problems

import (
	"sort"
	"strconv"
)

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func isDigitorialPermutation(n int) bool {
	s := 0
	for _, d := range strconv.Itoa(n) {
		s += factorials(int(d - '0'))
	}

	a := []byte(strconv.Itoa(s))
	b := []byte(strconv.Itoa(n))
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(a) == string(b)
}

func factorials(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}

	return res
}
