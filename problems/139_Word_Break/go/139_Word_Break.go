package problems

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for r := 1; r <= len(s); r++ {
		for l := 0; l < len(s); l++ {
			if dp[l] && isContain(s[l:r], wordDict) {
				dp[r] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func isContain(word string, wordDict []string) bool {
	for _, w := range wordDict {
		if w == word {
			return true
		}
	}

	return false
}
