package problems

import "testing"

func Test_coinChange(t *testing.T) {
	cases := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{name: "example 1", coins: []int{1, 2, 5}, amount: 11, want: 3},
		{name: "example 2", coins: []int{2}, amount: 3, want: -1},
		{name: "example 3 - zero amount", coins: []int{1}, amount: 0, want: 0},
		{name: "single coin exact", coins: []int{1}, amount: 1, want: 1},
		{name: "single coin multiple", coins: []int{1}, amount: 5, want: 5},
		{name: "greedy fails", coins: []int{1, 3, 4}, amount: 6, want: 2},
		{name: "large amount", coins: []int{1, 5, 10, 25}, amount: 100, want: 4},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}
