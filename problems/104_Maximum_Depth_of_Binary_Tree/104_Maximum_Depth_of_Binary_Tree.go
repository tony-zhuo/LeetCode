package problems

import "leet-code/structure"

func maxDepth(root *structure.TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}
