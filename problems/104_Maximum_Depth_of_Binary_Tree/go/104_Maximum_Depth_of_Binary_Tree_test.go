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
		root: []*int{libs.IntPtr(3), libs.IntPtr(9), libs.IntPtr(20), nil, nil, libs.IntPtr(15), libs.IntPtr(7)},
		want: 3,
	},
	{
		name: "example 2",
		root: []*int{libs.IntPtr(1), nil, libs.IntPtr(2)},
		want: 2,
	},
	{
		name: "empty tree",
		root: []*int{},
		want: 0,
	},
	{
		name: "single node",
		root: []*int{libs.IntPtr(1)},
		want: 1,
	},
	{
		name: "left only",
		root: []*int{libs.IntPtr(1), libs.IntPtr(2)},
		want: 2,
	},
}

func Test_maxDepth(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.Slice2BinaryTree(tt.root)
			got := maxDepth(root)
			if got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
