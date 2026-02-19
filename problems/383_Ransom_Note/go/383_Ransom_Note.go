package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func canConstruct(ransomNote string, magazine string) bool {
	cnt := make([]int, 26)

	for i := 0; i < len(magazine); i++ {
		cnt[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		cnt[ransomNote[i]-'a']--
	}

	for _, v := range cnt {
		if v < 0 {
			return false
		}
	}

	return true
}
