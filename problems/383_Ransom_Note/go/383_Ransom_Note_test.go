package problems

import "testing"

var cases = []struct {
	name       string
	ransomNote string
	magazine   string
	want       bool
}{
	{name: "example 1", ransomNote: "a", magazine: "b", want: false},
	{name: "example 2", ransomNote: "aa", magazine: "ab", want: false},
	{name: "example 3", ransomNote: "aa", magazine: "aab", want: true},
	{name: "empty ransom note", ransomNote: "", magazine: "abc", want: true},
	{name: "both empty", ransomNote: "", magazine: "", want: true},
	{name: "exact match", ransomNote: "abc", magazine: "abc", want: true},
	{name: "same letters different counts", ransomNote: "aab", magazine: "aba", want: true},
	{name: "magazine too short", ransomNote: "abcd", magazine: "abc", want: false},
}

func Test_canConstruct(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := canConstruct(tt.ransomNote, tt.magazine)
			if got != tt.want {
				t.Errorf("canConstruct(%q, %q) = %v, want %v", tt.ransomNote, tt.magazine, got, tt.want)
			}
		})
	}
}
