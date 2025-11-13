package problems

import "strings"

func reverseVowels(s string) string {
	bytes := []byte(s)
	left := 0
	right := len(bytes) - 1

	for left < right {
		if !isVowel(bytes[left]) {
			left++
		} else if !isVowel(bytes[right]) {
			right--
		} else {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}
	return string(bytes)
}

func isVowel(c byte) bool {
	return strings.ContainsRune("aeiouAEIOU", rune(c))
}
