package problems

import "testing"

var cases = []struct {
	name   string
	nums   []int
	target int
	want   int
}{
	{name: "example 1 - found in right half", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
	{name: "example 2 - not found", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
	{name: "example 3 - single element not found", nums: []int{1}, target: 0, want: -1},
	{name: "single element found", nums: []int{1}, target: 1, want: 0},
	{name: "not rotated - found", nums: []int{1, 2, 3, 4, 5}, target: 3, want: 2},
	{name: "not rotated - not found", nums: []int{1, 2, 3, 4, 5}, target: 6, want: -1},
	{name: "target at left pivot", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 4, want: 0},
	{name: "target at right end", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 2, want: 6},
	{name: "two elements rotated - found first", nums: []int{3, 1}, target: 3, want: 0},
	{name: "two elements rotated - found second", nums: []int{3, 1}, target: 1, want: 1},
	{name: "rotated by one", nums: []int{5, 1, 2, 3, 4}, target: 1, want: 1},
}

func Test_search(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}