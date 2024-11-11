package problems

import (
	"leet-code/libs"
	"leet-code/structure"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root []*int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: []*int{libs.IntPtr(3), libs.IntPtr(1), libs.IntPtr(4), nil, libs.IntPtr(2)},
				k:    1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				root: []*int{libs.IntPtr(5), libs.IntPtr(3), libs.IntPtr(6), libs.IntPtr(2), libs.IntPtr(4), nil, nil, libs.IntPtr(1)},
				k:    3,
			},
			want: 3,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootTree := structure.Slice2BinaryTree(tt.args.root)
			if got := kthSmallest(rootTree, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
