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
		name: "example 1 - with duplicate",
		nums: []int{1, 2, 2},
		want: [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{2, 2},
			{1, 2, 2},
		},
	},
	{
		name: "example 2 - single element",
		nums: []int{0},
		want: [][]int{
			{},
			{0},
		},
	},
	{
		name: "no duplicates",
		nums: []int{1, 2, 3},
		want: [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		},
	},
	{
		name: "all duplicates",
		nums: []int{4, 4, 4},
		want: [][]int{
			{},
			{4},
			{4, 4},
			{4, 4, 4},
		},
	},
	{
		name: "unsorted input with duplicates",
		nums: []int{4, 1, 4},
		want: [][]int{
			{},
			{1},
			{4},
			{1, 4},
			{4, 4},
			{1, 4, 4},
		},
	},
	{
		name: "negative duplicates",
		nums: []int{-1, -1, 0},
		want: [][]int{
			{},
			{-1},
			{0},
			{-1, -1},
			{-1, 0},
			{-1, -1, 0},
		},
	},
}

func sortSubsets(subsets [][]int) {
	for _, s := range subsets {
		sort.Ints(s)
	}
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		for k := range subsets[i] {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})
}

func equalSubsets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortSubsets(a)
	sortSubsets(b)
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

func Test_subsetsWithDup(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			got := subsetsWithDup(nums)
			if !equalSubsets(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
