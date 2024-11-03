package problems

import "leet-code/structure"

func isSubtree(root *structure.TreeNode, subRoot *structure.TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if (root != nil && subRoot == nil) || root == nil && subRoot != nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(tree1 *structure.TreeNode, tree2 *structure.TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	} else if (tree1 != nil && tree2 == nil) || (tree1 == nil && tree2 != nil) {
		return false
	} else if tree1.Val != tree2.Val {
		return false
	}

	return isSameTree(tree1.Left, tree2.Left) && isSameTree(tree1.Right, tree2.Right)
}
