package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

// 2 O(n)
func lengthOfLongestSubstring_2On(s string) int {
	maxLen := 0
	left := 0
	tmpMap := make(map[byte]bool)
	for right := 0; right < len(s); right++ {
		for tmpMap[s[right]] {
			delete(tmpMap, s[left])
			left++
		}

		tmpMap[s[right]] = true
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

// O(n)
func lengthOfLongestSubstring_On(s string) int {
	maxLen := 0
	left := 0
	tmpMap := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		if idx, ok := tmpMap[s[right]]; ok && idx >= left {
			left = idx + 1
		}

		tmpMap[s[right]] = right
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
