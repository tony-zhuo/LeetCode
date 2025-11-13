package problems

import "strings"

func reverseWords(s string) string {
	sArr := strings.Split(s, " ")
	var result []string
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] != "" {
			result = append(result, sArr[i])
		}
	}

	return strings.Join(result, " ")
}
