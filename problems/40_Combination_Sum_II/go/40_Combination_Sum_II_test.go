package problems

import (
	"sort"
	"testing"
)

var cases = []struct {
	name       string
	candidates []int
	target     int
	want       [][]int
}{
	{
		name:       "example 1 - with duplicates",
		candidates: []int{10, 1, 2, 7, 6, 1, 5},
		target:     8,
		want: [][]int{
			{1, 1, 6},
			{1, 2, 5},
			{1, 7},
			{2, 6},
		},
	},
	{
		name:       "example 2 - many duplicates",
		candidates: []int{2, 5, 2, 1, 2},
		target:     5,
		want: [][]int{
			{1, 2, 2},
			{5},
		},
	},
	{
		name:       "no solution",
		candidates: []int{2, 4, 6},
		target:     3,
		want:       [][]int{},
	},
	{
		name:       "single candidate equals target",
		candidates: []int{5},
		target:     5,
		want: [][]int{
			{5},
		},
	},
	{
		name:       "single candidate not equal target",
		candidates: []int{3},
		target:     5,
		want:       [][]int{},
	},
	{
		name:       "all duplicates contribute one combination",
		candidates: []int{1, 1, 1, 1},
		target:     3,
		want: [][]int{
			{1, 1, 1},
		},
	},
}

func sortCombos(combos [][]int) {
	for _, c := range combos {
		sort.Ints(c)
	}
	sort.Slice(combos, func(i, j int) bool {
		a, b := combos[i], combos[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
}

func equalCombos(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortCombos(a)
	sortCombos(b)
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

func Test_combinationSum2(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			candidates := make([]int, len(tt.candidates))
			copy(candidates, tt.candidates)
			got := combinationSum2(candidates, tt.target)
			if !equalCombos(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
