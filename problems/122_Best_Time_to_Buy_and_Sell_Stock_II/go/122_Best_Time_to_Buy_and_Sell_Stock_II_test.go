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
			want: 7,
		},
		{
			name: "example 2 - [1,2,3,4,5]",
			args: struct {
				prices []int
			}{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "example 3 - [7,6,4,3,1]",
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
			name: "simple up and down pattern",
			args: struct {
				prices []int
			}{prices: []int{1, 3, 2, 4}},
			want: 4,
		},
		{
			name: "multiple small gains",
			args: struct {
				prices []int
			}{prices: []int{1, 2, 1, 2, 1, 2}},
			want: 3,
		},
		{
			name: "large single transaction opportunity",
			args: struct {
				prices []int
			}{prices: []int{1, 100, 2, 200}},
			want: 297,
		},
		{
			name: "alternating pattern",
			args: struct {
				prices []int
			}{prices: []int{5, 1, 6, 2, 7, 3}},
			want: 10,
		},
		{
			name: "peak at beginning",
			args: struct {
				prices []int
			}{prices: []int{10, 1, 2, 3, 4}},
			want: 3,
		},
		{
			name: "valley at beginning",
			args: struct {
				prices []int
			}{prices: []int{1, 10, 9, 8, 7}},
			want: 9,
		},
		{
			name: "no transactions needed",
			args: struct {
				prices []int
			}{prices: []int{5, 4, 3, 2, 1}},
			want: 0,
		},
		{
			name: "one profitable day at end",
			args: struct {
				prices []int
			}{prices: []int{5, 4, 3, 2, 10}},
			want: 8,
		},
		{
			name: "multiple consecutive increases",
			args: struct {
				prices []int
			}{prices: []int{1, 3, 5, 7, 9}},
			want: 8,
		},
		{
			name: "buy low sell high multiple times",
			args: struct {
				prices []int
			}{prices: []int{2, 1, 4, 9}},
			want: 8,
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
