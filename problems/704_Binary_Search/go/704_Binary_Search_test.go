package problems

import "testing"

var cases = []struct {
	name   string
	nums   []int
	target int
	want   int
}{
	{name: "found in middle", nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
	{name: "not found", nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
	{name: "single element found", nums: []int{5}, target: 5, want: 0},
	{name: "single element not found", nums: []int{5}, target: -5, want: -1},
	{name: "two elements target last", nums: []int{2, 5}, target: 5, want: 1},
	{name: "two elements target first", nums: []int{2, 5}, target: 2, want: 0},
	{name: "target at first", nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 1, want: 0},
	{name: "target at last", nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 9, want: 8},
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