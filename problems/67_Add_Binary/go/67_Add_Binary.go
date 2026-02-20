package problems

import (
	"fmt"
	"strconv"
)

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || carry != 0 {
		total := carry
		if i >= 0 {
			total += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			total += int(b[j] - '0')
			j--
		}
		result = fmt.Sprintf("%s%s", strconv.Itoa(total%2), result)
		carry = total / 2
	}

	return result
}
