package problems

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smap := make(map[byte]int, len(s))
	tmap := make(map[byte]int, len(t))

	for i, _ := range s {
		smap[s[i]]++
		tmap[t[i]]++
	}

	for _, i := range s {
		tv, ok := tmap[byte(i)]
		if !ok {
			return false
		}
		if smap[byte(i)] != tv {
			return false
		}
	}

	return true
}
