package problems

import (
	"leet-code/structure"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *structure.TreeNode
		q *structure.TreeNode
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
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
