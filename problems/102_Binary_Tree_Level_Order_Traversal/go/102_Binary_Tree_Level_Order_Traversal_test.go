package problems

import (
	datastructures "leet-code/data_structures/binary_tree/go"
	"leet-code/libs"
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	root []*int
	want [][]int
}{
	{
		name: "example 1",
		root: []*int{libs.IntPtr(3), libs.IntPtr(9), libs.IntPtr(20), nil, nil, libs.IntPtr(15), libs.IntPtr(7)},
		want: [][]int{
			{3},
			{9, 20},
			{15, 7},
		},
	},
	{
		name: "single node",
		root: []*int{libs.IntPtr(1)},
		want: [][]int{
			{1},
		},
	},
	{
		name: "empty tree",
		root: []*int{},
		want: nil,
	},
}

func Test_levelOrder(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tree := datastructures.Slice2BinaryTree(tt.root)
			if got := levelOrder(tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrder_BFS(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tree := datastructures.Slice2BinaryTree(tt.root)
			got := levelOrder_BFS(tree)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder_BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}