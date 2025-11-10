package problems

func reverseVowels(s string) string {
	bytes := []byte(s)
	l := 0
	r := len(bytes) - 1

	for l < r {
		if !isVowel(bytes[l]) {
			l++
		} else if !isVowel(bytes[r]) {
			r--
		} else {
			// Both are vowels, swap them
			bytes[l], bytes[r] = bytes[r], bytes[l]
			l++
			r--
		}
	}
	return string(bytes)
}

func isVowel(c byte) bool {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	_, ok := vowels[c]
	return ok
}
