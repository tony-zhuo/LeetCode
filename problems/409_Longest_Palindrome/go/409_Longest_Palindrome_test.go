package problems

import "testing"

var cases = []struct {
	name string
	s    string
	want int
}{
	{name: "example 1", s: "abccccdd", want: 7},
	{name: "single char", s: "a", want: 1},
	{name: "all same pair", s: "aa", want: 2},
	{name: "odd count", s: "aaa", want: 3},
	{name: "already palindrome", s: "abcba", want: 5},
	{name: "all unique", s: "abcde", want: 1},
	{name: "case sensitive", s: "AaBb", want: 1},
	{name: "three same", s: "ccc", want: 3},
	{name: "mixed odds", s: "aaabbbccc", want: 7},
}

func Test_longestPalindrome_map(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome_map(tt.s); got != tt.want {
				t.Errorf("longestPalindrome_map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome_array(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome_array(tt.s); got != tt.want {
				t.Errorf("longestPalindrome_array() = %v, want %v", got, tt.want)
			}
		})
	}
}
