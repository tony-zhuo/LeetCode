package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap
func maximumXor(s string, t string) string {
	ones := 0
	for _, ch := range t {
		if ch == '1' {
			ones++
		}
	}

	zeros := len(t) - ones
	res := make([]byte, len(t))
	// XOR -> 相同為 0, 不同為 1
	for i := 0; i < len(t); i++ {
		if s[i] == '1' {
			if zeros > 0 {
				zeros--
				res[i] = '1'
			} else {
				ones--
				res[i] = '0'
			}
		} else {
			if ones > 0 {
				ones--
				res[i] = '1'
			} else {
				zeros--
				res[i] = '0'
			}
		}
	}

	return string(res)
}
