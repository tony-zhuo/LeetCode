package problems

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	mat  [][]int
	want [][]int
}{
	{
		name: "example 1",
		mat:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
	},
	{
		name: "example 2",
		mat:  [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
	},
	{
		name: "single zero",
		mat:  [][]int{{0}},
		want: [][]int{{0}},
	},
	{
		name: "single one",
		mat:  [][]int{{0, 1}},
		want: [][]int{{0, 1}},
	},
	{
		name: "large distance",
		mat:  [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}},
		want: [][]int{{4, 3, 2}, {3, 2, 1}, {2, 1, 0}},
	},
}

func Test_updateMatrix(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := updateMatrix(tt.mat)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}