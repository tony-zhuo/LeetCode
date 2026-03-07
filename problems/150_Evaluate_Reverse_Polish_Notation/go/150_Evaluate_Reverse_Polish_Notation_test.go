package problems

import "testing"

var cases = []struct {
	name   string
	tokens []string
	want   int
}{
	{
		name:   "example 1: simple addition",
		tokens: []string{"2", "1", "+", "3", "*"},
		want:   9,
	},
	{
		name:   "example 2: division",
		tokens: []string{"4", "13", "5", "/", "+"},
		want:   6,
	},
	{
		name:   "example 3: complex expression",
		tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		want:   22,
	},
	{
		name:   "single number",
		tokens: []string{"42"},
		want:   42,
	},
	{
		name:   "negative numbers",
		tokens: []string{"-1", "-2", "+"},
		want:   -3,
	},
}

func Test_evalRPN(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}