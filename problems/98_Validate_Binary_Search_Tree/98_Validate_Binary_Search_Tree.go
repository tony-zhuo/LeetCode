package problems

import (
	datastructures "leet-code/data_structures/binary_tree"
	"math"
)

func isValidBST(root *datastructures.TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *datastructures.TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	// 當前節點必須在 (min, max) 範圍內
	if node.Val <= min || node.Val >= max {
		return false
	}

	// 左子樹：所有節點必須 < node.Val（更新上界為 node.Val）
	// 右子樹：所有節點必須 > node.Val（更新下界為 node.Val）
	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}
