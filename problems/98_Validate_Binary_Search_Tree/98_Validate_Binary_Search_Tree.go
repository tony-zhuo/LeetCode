package problems

import "leet-code/structure"

func isValidBST(root *structure.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Val <= root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val >= root.Right.Val {
		return false
	}

	if !(checkChild(root.Left, root.Val, "left") && checkChild(root.Right, root.Val, "right")) {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

func checkChild(root *structure.TreeNode, val int, side string) bool {
	if root == nil {
		return true
	}

	switch side {
	case "left":
		if root.Val >= val {
			return false
		}
	case "right":
		if root.Val <= val {
			return false
		}
	}

	return checkChild(root.Left, val, side) && checkChild(root.Right, val, side)
}
