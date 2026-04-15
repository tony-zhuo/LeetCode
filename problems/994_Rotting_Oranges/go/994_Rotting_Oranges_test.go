package problems

import (
	"testing"
)

var cases = []struct {
	name string
	grid [][]int
	want int
}{
	{
		name: "example 1",
		grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
		want: 4,
	},
	{
		name: "example 2 unreachable fresh",
		grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
		want: -1,
	},
	{
		name: "example 3 no fresh",
		grid: [][]int{{0, 2}},
		want: 0,
	},
	{
		name: "all fresh no rotten",
		grid: [][]int{{1, 1}, {1, 1}},
		want: -1,
	},
	{
		name: "single rotten",
		grid: [][]int{{2}},
		want: 0,
	},
	{
		name: "all empty",
		grid: [][]int{{0, 0}, {0, 0}},
		want: 0,
	},
}

func cloneGrid(g [][]int) [][]int {
	out := make([][]int, len(g))
	for i, row := range g {
		out[i] = append([]int(nil), row...)
	}
	return out
}

func Test_orangesRotting(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := orangesRotting(cloneGrid(tt.grid))
			if got != tt.want {
				t.Errorf("orangesRotting() = %d, want %d", got, tt.want)
			}
		})
	}
}