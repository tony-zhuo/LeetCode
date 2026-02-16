package problems

func countSubstrings(s string) int {
	cnt := 0

	for i := 0; i < len(s); i++ {
		cnt += palindromeStr(s, i, i) + palindromeStr(s, i, i+1)
	}

	return cnt
}

func palindromeStr(s string, l, r int) int {
	cnt := 0
	for l >= 0 && r < len(s) && l <= r && s[l] == s[r] {
		cnt++
		l--
		r++
	}

	return cnt
}
