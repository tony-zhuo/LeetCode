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
		nums: []int{-1, 0, 1, 2, -1, -4},
		want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		name: "example 2",
		nums: []int{0, 1, 1},
		want: [][]int{},
	},
	{
		name: "example 3",
		nums: []int{0, 0, 0},
		want: [][]int{{0, 0, 0}},
	},
	{
		name: "no triplet",
		nums: []int{1, 2, 3},
		want: [][]int{},
	},
	{
		name: "multiple zeros",
		nums: []int{0, 0, 0, 0},
		want: [][]int{{0, 0, 0}},
	},
	{
		name: "all negatives",
		nums: []int{-1, -2, -3},
		want: [][]int{},
	},
}

func sortTriplets(triplets [][]int) {
	for _, t := range triplets {
		sort.Ints(t)
	}
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}

func equalTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortTriplets(a)
	sortTriplets(b)
	for i := range a {
		for k := 0; k < 3; k++ {
			if a[i][k] != b[i][k] {
				return false
			}
		}
	}
	return true
}

func Test_threeSum(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			got := threeSum(nums)
			if !equalTriplets(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
