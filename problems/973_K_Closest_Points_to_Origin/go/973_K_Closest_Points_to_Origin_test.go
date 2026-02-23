package problems

import (
	"sort"
	"testing"
)

var cases = []struct {
	name   string
	points [][]int
	k      int
	want   [][]int
}{
	{
		name:   "example 1",
		points: [][]int{{1, 3}, {-2, 2}},
		k:      1,
		want:   [][]int{{-2, 2}},
	},
	{
		name:   "example 2",
		points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
		k:      2,
		want:   [][]int{{3, 3}, {-2, 4}},
	},
	{
		name:   "k equals length",
		points: [][]int{{1, 1}, {2, 2}},
		k:      2,
		want:   [][]int{{1, 1}, {2, 2}},
	},
	{
		name:   "single point",
		points: [][]int{{0, 1}},
		k:      1,
		want:   [][]int{{0, 1}},
	},
	{
		name:   "origin included",
		points: [][]int{{0, 0}, {1, 1}, {2, 2}},
		k:      1,
		want:   [][]int{{0, 0}},
	},
}

// sortPoints sorts points for order-independent comparison.
func sortPoints(pts [][]int) {
	sort.Slice(pts, func(i, j int) bool {
		if pts[i][0] != pts[j][0] {
			return pts[i][0] < pts[j][0]
		}
		return pts[i][1] < pts[j][1]
	})
}

func equalPoints(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortPoints(a)
	sortPoints(b)
	for i := range a {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}

func Test_kClosest(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// deep copy to avoid mutating test data
			pts := make([][]int, len(tt.points))
			for i, p := range tt.points {
				pts[i] = []int{p[0], p[1]}
			}
			got := kClosest(pts, tt.k)
			if !equalPoints(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}