package problems

import (
	"leet-code/structure"
	"testing"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *structure.TreeNode
		subRoot *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
