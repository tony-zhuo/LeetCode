package problems

import "testing"

var cases = []struct {
	name string
	n    int
	want bool
}{
	{name: "1! = 1", n: 1, want: true},
	{name: "2! = 2", n: 2, want: true},
	{name: "factorion 145", n: 145, want: true},
	{name: "permutation of factorion 514", n: 514, want: true},
	{name: "3! = 6 not permutation of 3", n: 3, want: false},
	{name: "1!+0! = 2 not permutation of 10", n: 10, want: false},
}

func Test_isDigitorialPermutation(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := isDigitorialPermutation(tt.n)
			if got != tt.want {
				t.Errorf("isDigitorialPermutation(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
