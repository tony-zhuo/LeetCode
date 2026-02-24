package problems

import "testing"

var cases = []struct {
	name string
	s    string
	want int
}{
	{name: "example 1", s: "abcabcbb", want: 3},
	{name: "example 2", s: "bbbbb", want: 1},
	{name: "example 3", s: "pwwkew", want: 3},
	{name: "empty string", s: "", want: 0},
	{name: "single char", s: "a", want: 1},
	{name: "all unique", s: "abcdef", want: 6},
	{name: "with spaces", s: " ", want: 1},
	{name: "repeat at end", s: "abba", want: 2},
}

func Test_lengthOfLongestSubstring_2On(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring_2On(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring_2On(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring_On(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring_On(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring_On(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
