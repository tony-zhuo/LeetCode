package problems

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	word1Pointer := 0
	word2Pointer := 0
	var res strings.Builder // 使用 strings.Builder 來高效累加字串

	for word1Pointer < len(word1) || word2Pointer < len(word2) {
		if word1Pointer < len(word1) {
			res.WriteByte(word1[word1Pointer])
			word1Pointer++
		}

		if word2Pointer < len(word2) {
			res.WriteByte(word2[word2Pointer])
			word2Pointer++
		}
	}

	return res.String()
}
