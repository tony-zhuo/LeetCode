package problems

import (
	"leet-code/libs"
	datastructures "leet-code/data_structures/binary_tree/go"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root []*int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				root: []*int{libs.IntPtr(2), libs.IntPtr(1), libs.IntPtr(3)},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				root: []*int{libs.IntPtr(5), libs.IntPtr(1), libs.IntPtr(4), nil, nil, libs.IntPtr(3), libs.IntPtr(6)},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				root: []*int{libs.IntPtr(5), libs.IntPtr(4), libs.IntPtr(6), nil, nil, libs.IntPtr(3), libs.IntPtr(7)},
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := datastructures.Slice2BinaryTree(tt.args.root)
			if got := isValidBST(tree); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
