package problems

import (
	"fmt"
	"strconv"
)

type Codec struct{}

func (c *Codec) Encode(strs []string) string {
	res := ""
	for _, str := range strs {
		res += fmt.Sprintf("%d#%s", len(str), str)
	}

	return res
}

func (c *Codec) Decode(s string) []string {
	res := make([]string, 0)

	i := 0
	for i < len(s) {
		if string(s[i]) == "#" {
			length, _ := strconv.Atoi(string(s[i-1]))
			tmp := s[i+1 : i+1+length]
			res = append(res, tmp)
		}
		i++
	}

	return res
}

/*
"neet", "code", "love", "you"
4#neet4#code4#love4#you


*/
