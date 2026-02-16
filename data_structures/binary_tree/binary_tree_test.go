package datastructures

import (
	"reflect"
	"testing"
)

// Helper to build a balanced tree:
//
//	    1
//	   / \
//	  2   3
//	 / \
//	4   5
func buildBalancedTree() *TreeNode {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	return root
}

// Helper to build a left-skewed tree:
//
//	1
//	 \
//	  (nil)
//	1 -> 2 -> 3
func buildLeftSkewedTree() *TreeNode {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(3)
	return root
}

// Helper to build a right-skewed tree:
//
//	1 -> 2 -> 3
func buildRightSkewedTree() *TreeNode {
	root := NewTreeNode(1)
	root.Right = NewTreeNode(2)
	root.Right.Right = NewTreeNode(3)
	return root
}

func TestNewTreeNode(t *testing.T) {
	node := NewTreeNode(42)
	if node.Val != 42 {
		t.Errorf("expected Val=42, got %d", node.Val)
	}
	if node.Left != nil {
		t.Error("expected Left to be nil")
	}
	if node.Right != nil {
		t.Error("expected Right to be nil")
	}
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "nil root",
			root: nil,
			want: []int{},
		},
		{
			name: "single node",
			root: NewTreeNode(1),
			want: []int{1},
		},
		{
			name: "balanced tree",
			root: buildBalancedTree(),
			want: []int{4, 2, 5, 1, 3},
		},
		{
			name: "left-skewed tree",
			root: buildLeftSkewedTree(),
			want: []int{3, 2, 1},
		},
		{
			name: "right-skewed tree",
			root: buildRightSkewedTree(),
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderTraversal(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "nil root",
			root: nil,
			want: []int{},
		},
		{
			name: "single node",
			root: NewTreeNode(1),
			want: []int{1},
		},
		{
			name: "balanced tree",
			root: buildBalancedTree(),
			want: []int{1, 2, 4, 5, 3},
		},
		{
			name: "left-skewed tree",
			root: buildLeftSkewedTree(),
			want: []int{1, 2, 3},
		},
		{
			name: "right-skewed tree",
			root: buildRightSkewedTree(),
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreorderTraversal(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "nil root",
			root: nil,
			want: []int{},
		},
		{
			name: "single node",
			root: NewTreeNode(1),
			want: []int{1},
		},
		{
			name: "balanced tree",
			root: buildBalancedTree(),
			want: []int{4, 5, 2, 3, 1},
		},
		{
			name: "left-skewed tree",
			root: buildLeftSkewedTree(),
			want: []int{3, 2, 1},
		},
		{
			name: "right-skewed tree",
			root: buildRightSkewedTree(),
			want: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PostorderTraversal(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "nil root",
			root: nil,
			want: [][]int{},
		},
		{
			name: "single node",
			root: NewTreeNode(1),
			want: [][]int{{1}},
		},
		{
			name: "balanced tree",
			root: buildBalancedTree(),
			want: [][]int{{1}, {2, 3}, {4, 5}},
		},
		{
			name: "left-skewed tree",
			root: buildLeftSkewedTree(),
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "right-skewed tree",
			root: buildRightSkewedTree(),
			want: [][]int{{1}, {2}, {3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LevelOrderTraversal(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeight(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "nil root",
			root: nil,
			want: 0,
		},
		{
			name: "single node",
			root: NewTreeNode(1),
			want: 1,
		},
		{
			name: "balanced tree",
			root: buildBalancedTree(),
			want: 3,
		},
		{
			name: "left-skewed tree",
			root: buildLeftSkewedTree(),
			want: 3,
		},
		{
			name: "right-skewed tree",
			root: buildRightSkewedTree(),
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Height(tt.root)
			if got != tt.want {
				t.Errorf("Height() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCountNodes(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "nil root",
			root: nil,
			want: 0,
		},
		{
			name: "single node",
			root: NewTreeNode(1),
			want: 1,
		},
		{
			name: "balanced tree",
			root: buildBalancedTree(),
			want: 5,
		},
		{
			name: "left-skewed tree",
			root: buildLeftSkewedTree(),
			want: 3,
		},
		{
			name: "right-skewed tree",
			root: buildRightSkewedTree(),
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountNodes(tt.root)
			if got != tt.want {
				t.Errorf("CountNodes() = %d, want %d", got, tt.want)
			}
		})
	}
}
