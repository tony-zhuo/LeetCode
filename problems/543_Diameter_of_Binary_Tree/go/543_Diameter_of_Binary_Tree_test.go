package problems

import (
	"testing"

	datastructures "leet-code/data_structures/binary_tree/go"
	"leet-code/libs"
)

var cases = []struct {
	name string
	root []*int
	want int
}{
	{
		name: "example 1",
		root: []*int{libs.IntPtr(1), libs.IntPtr(2), libs.IntPtr(3), libs.IntPtr(4), libs.IntPtr(5)},
		want: 3,
	},
	{
		name: "example 2",
		root: []*int{libs.IntPtr(1), libs.IntPtr(2)},
		want: 1,
	},
	{
		name: "single node",
		root: []*int{libs.IntPtr(1)},
		want: 0,
	},
	{
		name: "longest path not through root",
		root: []*int{libs.IntPtr(1), libs.IntPtr(2), nil, libs.IntPtr(3), libs.IntPtr(4), libs.IntPtr(5), nil, nil, nil, libs.IntPtr(6)},
		want: 4,
	},
}

func Test_diameterOfBinaryTree(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.Slice2BinaryTree(tt.root)
			got := diameterOfBinaryTree(root)
			if got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
