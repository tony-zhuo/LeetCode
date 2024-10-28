package problems

import (
	"leet-code/libs"
	"leet-code/structure"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p []*int
		q []*int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				p: []*int{libs.IntPtr(1), libs.IntPtr(2), libs.IntPtr(3)},
				q: []*int{libs.IntPtr(1), libs.IntPtr(2), libs.IntPtr(3)},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				p: []*int{libs.IntPtr(1), libs.IntPtr(2)},
				q: []*int{libs.IntPtr(1), nil, libs.IntPtr(2)},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				p: []*int{libs.IntPtr(1), libs.IntPtr(2), libs.IntPtr(1)},
				q: []*int{libs.IntPtr(1), libs.IntPtr(1), libs.IntPtr(2)},
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pTree := structure.Slice2BinaryTree(tt.args.p)
			qTree := structure.Slice2BinaryTree(tt.args.q)
			if got := isSameTree(pTree, qTree); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
