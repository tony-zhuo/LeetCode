package problems

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "classic palindrome", s: "A man, a plan, a canal: Panama", want: true},
		{name: "not palindrome", s: "race a car", want: false},
		{name: "single space", s: " ", want: true},
		{name: "empty string", s: "", want: true},
		{name: "single char", s: "a", want: true},
		{name: "two different chars", s: "ab", want: false},
		{name: "odd palindrome", s: "aba", want: true},
		{name: "alphanumeric mix", s: "0P", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
