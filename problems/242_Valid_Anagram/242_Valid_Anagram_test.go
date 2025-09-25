package problems

import "testing"

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			s string
			t string
		}
		want bool
	}{
		{
			name: "valid anagram - anagram/nagaram",
			args: struct {
				s string
				t string
			}{s: "anagram", t: "nagaram"},
			want: true,
		},
		{
			name: "invalid anagram - rat/car",
			args: struct {
				s string
				t string
			}{s: "rat", t: "car"},
			want: false,
		},
		{
			name: "empty strings",
			args: struct {
				s string
				t string
			}{s: "", t: ""},
			want: true,
		},
		{
			name: "single character match",
			args: struct {
				s string
				t string
			}{s: "a", t: "a"},
			want: true,
		},
		{
			name: "single character different",
			args: struct {
				s string
				t string
			}{s: "a", t: "b"},
			want: false,
		},
		{
			name: "different lengths",
			args: struct {
				s string
				t string
			}{s: "abc", t: "ab"},
			want: false,
		},
		{
			name: "same characters different frequency",
			args: struct {
				s string
				t string
			}{s: "aab", t: "abb"},
			want: false,
		},
		{
			name: "valid anagram with repeated characters",
			args: struct {
				s string
				t string
			}{s: "listen", t: "silent"},
			want: true,
		},
		{
			name: "case sensitive - different cases",
			args: struct {
				s string
				t string
			}{s: "A", t: "a"},
			want: false,
		},
		{
			name: "longer valid anagram",
			args: struct {
				s string
				t string
			}{s: "conversation", t: "voices rant on"},
			want: false,
		},
		{
			name: "palindrome anagram",
			args: struct {
				s string
				t string
			}{s: "evil", t: "vile"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
