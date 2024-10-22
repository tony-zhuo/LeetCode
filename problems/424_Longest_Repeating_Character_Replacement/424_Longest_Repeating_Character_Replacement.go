package problems

func characterReplacement(s string, k int) int {
	l, r := 0, 0
	mapCount := make(map[byte]int)
	res := 0
	maxCount := 0

	for r < len(s) {
		mapCount[s[r]]++
		length := r - l + 1
		maxCount = max(maxCount, mapCount[s[r]])
		if length-maxCount > k {
			mapCount[s[l]]--
			l++
		} else {
			res = max(res, length)
		}

		r++
	}

	return res
}

/*
ABAB
k = 2

A BAB
l = 0
r = 0
length = 1
max count = 1

AB AB
l = 0
r = 1
length = 2
max count = 1

ABA B
l = 0
r = 2
length = 3
max count = 2
3 - 2 <= k

ABAB
l = 0
r = 3
length = 4
max count = 2
4 - 2 <= k
*/
