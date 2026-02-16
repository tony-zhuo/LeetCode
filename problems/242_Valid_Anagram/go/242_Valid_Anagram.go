package problems

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[rune]int)

	for _, ch := range s {
		counter[ch]++
	}

	for _, ch := range t {
		counter[ch]--
		if counter[ch] < 0 {
			return false
		}
	}

	return true
}
