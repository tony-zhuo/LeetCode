package problems

import "testing"

var cases = []struct {
	name string
	nums []int
	want int
}{
	{name: "example 1", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
	{name: "example 2 - single element", nums: []int{1}, want: 1},
	{name: "example 3 - all positive prefix", nums: []int{5, 4, -1, 7, 8}, want: 23},
	{name: "all negative", nums: []int{-3, -2, -5, -1}, want: -1},
	{name: "two elements", nums: []int{-1, 2}, want: 2},
}

func Test_maxSubArray(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("maxSubArray(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
