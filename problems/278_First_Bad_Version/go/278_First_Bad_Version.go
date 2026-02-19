package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap
// isBadVersion is the API provided by the problem.
// It is defined as a variable so tests can override it.
var isBadVersion func(version int) bool

func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
