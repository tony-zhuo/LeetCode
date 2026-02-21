package problems

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name        string
	intervals   [][]int
	newInterval []int
	want        [][]int
}{
	{
		name:        "example 1 - no overlap",
		intervals:   [][]int{{1, 3}, {6, 9}},
		newInterval: []int{2, 5},
		want:        [][]int{{1, 5}, {6, 9}},
	},
	{
		name:        "example 2 - multiple merges",
		intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
		newInterval: []int{4, 8},
		want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
	},
	{
		name:        "empty intervals",
		intervals:   [][]int{},
		newInterval: []int{5, 7},
		want:        [][]int{{5, 7}},
	},
	{
		name:        "insert before all",
		intervals:   [][]int{{3, 5}, {6, 9}},
		newInterval: []int{1, 2},
		want:        [][]int{{1, 2}, {3, 5}, {6, 9}},
	},
	{
		name:        "insert after all",
		intervals:   [][]int{{1, 2}, {3, 5}},
		newInterval: []int{6, 8},
		want:        [][]int{{1, 2}, {3, 5}, {6, 8}},
	},
}

func Test_insert(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, got, tt.want)
			}
		})
	}
}
