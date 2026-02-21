package problems

import "testing"

var cases = []struct {
	name string
	nums []int
	want bool
}{
	{name: "example 1 - has duplicate", nums: []int{1, 2, 3, 1}, want: true},
	{name: "example 2 - no duplicate", nums: []int{1, 2, 3, 4}, want: false},
	{name: "example 3 - multiple duplicates", nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	{name: "single element", nums: []int{1}, want: false},
	{name: "two same", nums: []int{1, 1}, want: true},
}

func Test_containsDuplicate(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
