package problems

func characterReplacement(s string, k int) int {
	l := 0
	maxCount := 0
	res := 0
	m := make(map[byte]int, len(s))
	for r := 0; r < len(s); r++ {
		m[s[r]]++
		stringLen := r - l + 1
		maxCount = max(m[s[r]], maxCount)
		if stringLen-maxCount > k {
			m[s[l]]--
			l++
		} else {
			res = max(res, stringLen)
		}
	}

	return res
}

/*
A ABABBA
string len 1
max len 1
k = 1
l dont move

AA BABBA
string len 2
max len 2
k = 1


AAB ABBA
string len 3
max len 2
k = 1
l 不用動

AABA BBA
string len 4
max len 3
k = 1
string len - max len <= k，l 不用動

AABAB BA
string len 5
max len 3
k = 1
string len - max len > k，l 往右一格

A ABABB A
string len 5
max len 3（B 有三個
k = 1
string len - max len > k，l 往右一格

AA BABBA
string len 5
max len 3（B 有三個
k = 1
string len - max len > k，l 往右一格
*/
