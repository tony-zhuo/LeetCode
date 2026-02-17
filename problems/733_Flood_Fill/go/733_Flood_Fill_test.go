package problems

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name  string
	image [][]int
	sr    int
	sc    int
	color int
	want  [][]int
}{
	{
		name:  "example 1",
		image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		sr: 1, sc: 1, color: 2,
		want: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},
	{
		name:  "same color no change",
		image: [][]int{{0, 0, 0}, {0, 0, 0}},
		sr: 0, sc: 0, color: 0,
		want: [][]int{{0, 0, 0}, {0, 0, 0}},
	},
	{
		name:  "fill entire grid",
		image: [][]int{{0, 0, 0}, {0, 0, 0}},
		sr: 0, sc: 0, color: 2,
		want: [][]int{{2, 2, 2}, {2, 2, 2}},
	},
	{
		name:  "single cell",
		image: [][]int{{1}},
		sr: 0, sc: 0, color: 2,
		want: [][]int{{2}},
	},
	{
		name:  "cross shape",
		image: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
		sr: 1, sc: 1, color: 3,
		want: [][]int{{0, 3, 0}, {3, 3, 3}, {0, 3, 0}},
	},
	{
		name:  "diamond shape",
		image: [][]int{{1, 2, 1}, {2, 2, 2}, {1, 2, 1}},
		sr: 1, sc: 1, color: 5,
		want: [][]int{{1, 5, 1}, {5, 5, 5}, {1, 5, 1}},
	},
}

func copyImage(img [][]int) [][]int {
	cp := make([][]int, len(img))
	for i := range img {
		cp[i] = make([]int, len(img[i]))
		copy(cp[i], img[i])
	}
	return cp
}

func Test_floodFill(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := floodFill(copyImage(tt.image), tt.sr, tt.sc, tt.color)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
