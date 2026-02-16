package problems

import (
	"leet-code/libs"
	datastructures "leet-code/data_structures/binary_tree"
	"testing"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    []*int
		subRoot []*int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				root:    []*int{libs.IntPtr(3), libs.IntPtr(4), libs.IntPtr(5), libs.IntPtr(1), libs.IntPtr(2)},
				subRoot: []*int{libs.IntPtr(4), libs.IntPtr(1), libs.IntPtr(2)},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				root:    []*int{libs.IntPtr(3), libs.IntPtr(4), libs.IntPtr(5), libs.IntPtr(1), libs.IntPtr(2), nil, nil, nil, nil, libs.IntPtr(0)},
				subRoot: []*int{libs.IntPtr(4), libs.IntPtr(1), libs.IntPtr(2)},
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootTree := datastructures.Slice2BinaryTree(tt.args.root)
			subRootTree := datastructures.Slice2BinaryTree(tt.args.subRoot)
			if got := isSubtree(rootTree, subRootTree); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
