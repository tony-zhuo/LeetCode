package problems

import "testing"

var cases = []struct {
	name string
	nums []int
	want int
}{
	{name: "all even - player 1 takes all", nums: []int{4, 2}, want: 6},
	{name: "single odd - switches to player 2", nums: []int{1}, want: -1},
	{name: "odd switches back and forth", nums: []int{3, 1, 4}, want: 2},
	{name: "mixed odd even", nums: []int{1, 2, 3}, want: 0},
	{name: "6th round switch", nums: []int{2, 2, 2, 2, 2, 2}, want: 8},
	{name: "6th round and odd cancel out", nums: []int{2, 2, 2, 2, 2, 3}, want: 13},
}

func Test_scoreDifference(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := scoreDifference(tt.nums)
			if got != tt.want {
				t.Errorf("scoreDifference(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
