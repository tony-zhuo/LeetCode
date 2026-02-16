package problems

import (
	"leet-code/libs"
	datastructures "leet-code/data_structures/binary_tree/go"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root []*int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				root: []*int{libs.IntPtr(3), libs.IntPtr(9), libs.IntPtr(20), nil, nil, libs.IntPtr(15), libs.IntPtr(7)},
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 17},
			},
		},
		{
			name: "",
			args: args{
				root: []*int{libs.IntPtr(1)},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "",
			args: args{
				root: []*int{},
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := datastructures.Slice2BinaryTree(tt.args.root)
			if got := levelOrder(tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
