package problems

import "testing"

var cases = []struct {
	name string
	nums []int
	want int
}{
	{name: "example 1", nums: []int{3, 2, 3}, want: 3},
	{name: "example 2", nums: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
	{name: "single element", nums: []int{1}, want: 1},
	{name: "two same", nums: []int{6, 5, 5}, want: 5},
}

func Test_majorityElement(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("majorityElement(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
