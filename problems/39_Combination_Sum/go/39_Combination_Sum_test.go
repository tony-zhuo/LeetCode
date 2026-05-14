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
		name:       "example 1",
		candidates: []int{2, 3, 6, 7},
		target:     7,
		want: [][]int{
			{2, 2, 3},
			{7},
		},
	},
	{
		name:       "example 2 - multiple combinations",
		candidates: []int{2, 3, 5},
		target:     8,
		want: [][]int{
			{2, 2, 2, 2},
			{2, 3, 3},
			{3, 5},
		},
	},
	{
		name:       "example 3 - no solution",
		candidates: []int{2},
		target:     1,
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
		name:       "unsorted candidates",
		candidates: []int{8, 7, 4, 3},
		target:     11,
		want: [][]int{
			{3, 4, 4},
			{3, 8},
			{4, 7},
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

func Test_combinationSum(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			candidates := make([]int, len(tt.candidates))
			copy(candidates, tt.candidates)
			got := combinationSum(candidates, tt.target)
			if !equalCombos(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
