package problems

import (
	"leet-code/libs"
	"leet-code/structure"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *structure.TreeNode
	}{
		{
			name: "",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: structure.Slice2BinaryTree([]*int{
				libs.IntPtr(3), libs.IntPtr(9), libs.IntPtr(20), nil, nil, libs.IntPtr(15), libs.IntPtr(7),
			}),
		},
		{
			name: "",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: structure.Slice2BinaryTree([]*int{
				libs.IntPtr(-1),
			}),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
