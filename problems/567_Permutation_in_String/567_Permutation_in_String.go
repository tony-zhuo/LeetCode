package problems

/*
use sliding window

for example
s1: ab
s2: eidbaooo
*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Cnt, windowCnt [26]int

	for i := 0; i < len(s1); i++ {
		s1Cnt[s1[i]-'a']++
		windowCnt[s2[i]-'a']++
	}
	if s1Cnt == windowCnt {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		windowCnt[s2[i]-'a']++
		windowCnt[s2[i-len(s1)]-'a']--

		if s1Cnt == windowCnt {
			return true
		}
	}

	return false
}
