package problems

import (
	"testing"

	datastructures "leet-code/data_structures/binary_tree/go"
	"leet-code/libs"
)

// findNode returns the node with the given value in the tree.
func findNode(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

var cases = []struct {
	name string
	root []*int
	p    int
	q    int
	want int
}{
	{
		name: "LCA is root",
		root: []*int{libs.IntPtr(6), libs.IntPtr(2), libs.IntPtr(8), libs.IntPtr(0), libs.IntPtr(4), libs.IntPtr(7), libs.IntPtr(9), nil, nil, libs.IntPtr(3), libs.IntPtr(5)},
		p:    2,
		q:    8,
		want: 6,
	},
	{
		name: "LCA is one of the nodes",
		root: []*int{libs.IntPtr(6), libs.IntPtr(2), libs.IntPtr(8), libs.IntPtr(0), libs.IntPtr(4), libs.IntPtr(7), libs.IntPtr(9), nil, nil, libs.IntPtr(3), libs.IntPtr(5)},
		p:    2,
		q:    4,
		want: 2,
	},
	{
		name: "two nodes",
		root: []*int{libs.IntPtr(2), libs.IntPtr(1)},
		p:    2,
		q:    1,
		want: 2,
	},
	{
		name: "both in right subtree",
		root: []*int{libs.IntPtr(6), libs.IntPtr(2), libs.IntPtr(8), libs.IntPtr(0), libs.IntPtr(4), libs.IntPtr(7), libs.IntPtr(9)},
		p:    7,
		q:    9,
		want: 8,
	},
	{
		name: "same node",
		root: []*int{libs.IntPtr(6), libs.IntPtr(2), libs.IntPtr(8)},
		p:    2,
		q:    2,
		want: 2,
	},
}

func Test_lowestCommonAncestor(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.Slice2BinaryTree(tt.root)
			p := findNode(root, tt.p)
			q := findNode(root, tt.q)
			got := lowestCommonAncestor(root, p, q)
			if got == nil || got.Val != tt.want {
				gotVal := -1
				if got != nil {
					gotVal = got.Val
				}
				t.Errorf("lowestCommonAncestor() = %v, want %v", gotVal, tt.want)
			}
		})
	}
}