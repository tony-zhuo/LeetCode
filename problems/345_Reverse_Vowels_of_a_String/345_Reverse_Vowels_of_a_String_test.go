package problems

import "testing"

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "hello",
			args: "hello",
			want: "holle",
		},
		{
			name: "leetcode",
			args: "leetcode",
			want: "leotcede",
		},
		{
			name: "aA",
			args: "aA",
			want: "Aa",
		},
		{
			name: "single vowel",
			args: "a",
			want: "a",
		},
		{
			name: "no vowels",
			args: "bcdfg",
			want: "bcdfg",
		},
		{
			name: "all vowels",
			args: "aeiou",
			want: "uoiea",
		},
		{
			name: "mixed case",
			args: "AeIoU",
			want: "UoIeA",
		},
		{
			name: "empty string",
			args: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
