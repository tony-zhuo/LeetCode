package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func longestPalindrome_map(s string) int {
	charMap := make(map[rune]int, 0)
	for _, item := range s {
		if _, found := charMap[item]; found {
			charMap[item]++
		} else {
			charMap[item] = 1
		}
	}

	res := 0
	hasOdd := false

	for _, cnt := range charMap {
		res += cnt / 2 * 2
		if cnt%2 == 1 {
			hasOdd = true
		}
	}

	if hasOdd {
		res += 1
	}

	return res
}

func longestPalindrome_array(s string) int {
	var freq [128]int
	for _, c := range s {
		freq[c]++
	}

	res := 0
	hasOdd := false

	for _, cnt := range freq {
		res += cnt / 2 * 2
		if cnt%2 == 1 {
			hasOdd = true
		}
	}

	if hasOdd {
		res += 1
	}

	return res
}
