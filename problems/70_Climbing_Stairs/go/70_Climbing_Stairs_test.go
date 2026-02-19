package problems

import "testing"

var cases = []struct {
	name string
	n    int
	want int
}{
	{name: "n=1", n: 1, want: 1},
	{name: "n=2", n: 2, want: 2},
	{name: "n=3", n: 3, want: 3},
	{name: "n=4", n: 4, want: 5},
	{name: "n=5", n: 5, want: 8},
	{name: "n=10", n: 10, want: 89},
	{name: "n=20", n: 20, want: 10946},
	{name: "n=45", n: 45, want: 1836311903},
}

func Test_climbStairs_hashmap(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs_hashmap(tt.n); got != tt.want {
				t.Errorf("climbStairs_hashmap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs_dp(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs_dp(tt.n); got != tt.want {
				t.Errorf("climbStairs_dp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs_matrix(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs_matrix(tt.n); got != tt.want {
				t.Errorf("climbStairs_matrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
