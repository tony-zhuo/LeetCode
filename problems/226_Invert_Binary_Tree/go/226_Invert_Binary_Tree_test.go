package problems

import (
	"reflect"
	"testing"

	datastructures "leet-code/data_structures/binary_tree/go"
	"leet-code/libs"
)

func Test_invertTree(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want []*int
	}{
		{
			name: "example 1",
			root: []*int{libs.IntPtr(4), libs.IntPtr(2), libs.IntPtr(7), libs.IntPtr(1), libs.IntPtr(3), libs.IntPtr(6), libs.IntPtr(9)},
			want: []*int{libs.IntPtr(4), libs.IntPtr(7), libs.IntPtr(2), libs.IntPtr(9), libs.IntPtr(6), libs.IntPtr(3), libs.IntPtr(1)},
		},
		{
			name: "example 2",
			root: []*int{libs.IntPtr(2), libs.IntPtr(1), libs.IntPtr(3)},
			want: []*int{libs.IntPtr(2), libs.IntPtr(3), libs.IntPtr(1)},
		},
		{
			name: "empty tree",
			root: []*int{},
			want: []*int{},
		},
		{
			name: "single node",
			root: []*int{libs.IntPtr(1)},
			want: []*int{libs.IntPtr(1)},
		},
		{
			name: "left only",
			root: []*int{libs.IntPtr(1), libs.IntPtr(2)},
			want: []*int{libs.IntPtr(1), nil, libs.IntPtr(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.Slice2BinaryTree(tt.root)
			wantTree := datastructures.Slice2BinaryTree(tt.want)
			got := invertTree(root)
			gotArr := tree2Slice(got)
			wantArr := tree2Slice(wantTree)
			if !reflect.DeepEqual(gotArr, wantArr) {
				t.Errorf("invertTree() = %v, want %v", gotArr, wantArr)
			}
		})
	}
}

func tree2Slice(root *datastructures.TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	result := []interface{}{}
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	// trim trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}
