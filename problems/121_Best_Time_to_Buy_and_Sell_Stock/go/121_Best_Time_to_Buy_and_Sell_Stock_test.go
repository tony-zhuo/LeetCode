package problems

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			prices []int
		}
		want int
	}{
		{
			name: "example 1 - [7,1,5,3,6,4]",
			args: struct {
				prices []int
			}{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "example 2 - [7,6,4,3,1]",
			args: struct {
				prices []int
			}{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "empty array",
			args: struct {
				prices []int
			}{prices: []int{}},
			want: 0,
		},
		{
			name: "single price",
			args: struct {
				prices []int
			}{prices: []int{5}},
			want: 0,
		},
		{
			name: "two prices - profit available",
			args: struct {
				prices []int
			}{prices: []int{1, 5}},
			want: 4,
		},
		{
			name: "two prices - no profit",
			args: struct {
				prices []int
			}{prices: []int{5, 1}},
			want: 0,
		},
		{
			name: "all same prices",
			args: struct {
				prices []int
			}{prices: []int{3, 3, 3, 3}},
			want: 0,
		},
		{
			name: "increasing prices",
			args: struct {
				prices []int
			}{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "decreasing prices",
			args: struct {
				prices []int
			}{prices: []int{5, 4, 3, 2, 1}},
			want: 0,
		},
		{
			name: "minimum at beginning, maximum at end",
			args: struct {
				prices []int
			}{prices: []int{1, 3, 2, 4, 1, 6}},
			want: 5,
		},
		{
			name: "multiple peaks and valleys",
			args: struct {
				prices []int
			}{prices: []int{3, 2, 6, 5, 0, 3}},
			want: 4,
		},
		{
			name: "large numbers",
			args: struct {
				prices []int
			}{prices: []int{10000, 1, 10001}},
			want: 10000,
		},
		{
			name: "profit at the end",
			args: struct {
				prices []int
			}{prices: []int{2, 4, 1, 7}},
			want: 6,
		},
		{
			name: "zigzag pattern",
			args: struct {
				prices []int
			}{prices: []int{1, 4, 2, 6, 3, 7}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
