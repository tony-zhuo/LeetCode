package problems

import (
	"sort"
	"testing"
)

var cases = []struct {
	name string
	nums []int
	want [][]int
}{
	{
		name: "example 1",
		nums: []int{1, 2, 3},
		want: [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		},
	},
	{
		name: "example 2 - two elements",
		nums: []int{0, 1},
		want: [][]int{
			{0, 1},
			{1, 0},
		},
	},
	{
		name: "example 3 - single element",
		nums: []int{1},
		want: [][]int{
			{1},
		},
	},
	{
		name: "negative numbers",
		nums: []int{-1, 0, 1},
		want: [][]int{
			{-1, 0, 1},
			{-1, 1, 0},
			{0, -1, 1},
			{0, 1, -1},
			{1, -1, 0},
			{1, 0, -1},
		},
	},
}

func sortPermutations(perms [][]int) {
	sort.Slice(perms, func(i, j int) bool {
		for k := range perms[i] {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}
		return false
	})
}

func equalPermutations(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortPermutations(a)
	sortPermutations(b)
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for k := range a[i] {
			if a[i][k] != b[i][k] {
				return false
			}
		}
	}
	return true
}

func Test_permute(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			got := permute(nums)
			if !equalPermutations(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
