package problems

import datastructures "leet-code/data_structures/binary_tree/go"

func maxDepth(root *datastructures.TreeNode) int {
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
