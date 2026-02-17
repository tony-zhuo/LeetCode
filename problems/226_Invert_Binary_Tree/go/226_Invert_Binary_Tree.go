package problems

import datastructures "leet-code/data_structures/binary_tree/go"

func invertTree(root *datastructures.TreeNode) *datastructures.TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)

	return root
}
