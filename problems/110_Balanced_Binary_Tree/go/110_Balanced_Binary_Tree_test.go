package problems

import (
	"testing"

	datastructures "leet-code/data_structures/binary_tree/go"
	"leet-code/libs"
)

var cases = []struct {
	name string
	root []*int
	want bool
}{
	{
		name: "balanced tree",
		root: []*int{libs.IntPtr(3), libs.IntPtr(9), libs.IntPtr(20), nil, nil, libs.IntPtr(15), libs.IntPtr(7)},
		want: true,
	},
	{
		name: "unbalanced tree",
		root: []*int{libs.IntPtr(1), libs.IntPtr(2), libs.IntPtr(2), libs.IntPtr(3), libs.IntPtr(3), nil, nil, libs.IntPtr(4), libs.IntPtr(4)},
		want: false,
	},
	{
		name: "empty tree",
		root: []*int{},
		want: true,
	},
	{
		name: "single node",
		root: []*int{libs.IntPtr(1)},
		want: true,
	},
	{
		name: "left skewed",
		root: []*int{libs.IntPtr(1), libs.IntPtr(2), nil, libs.IntPtr(3)},
		want: false,
	},
	{
		name: "full but unbalanced",
		root: []*int{libs.IntPtr(1), libs.IntPtr(2), libs.IntPtr(3), libs.IntPtr(4), libs.IntPtr(5), libs.IntPtr(6), nil, libs.IntPtr(8)},
		want: true,
	},
}

func Test_isBalanced(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.Slice2BinaryTree(tt.root)
			got := isBalanced(root)
			if got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
