package problems

func countSubstrings(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		l, r := i, i
		// 基數
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}

		l, r = i, i+1
		// 偶數
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}

	return res
}
