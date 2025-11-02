package problems

func gcdOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)
	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLen := gcb(len1, len2)
	return str1[:gcdLen]
}

func gcb(a, b int) int {
	if b == 0 {
		return a
	}
	return gcb(b, a%b)
}
