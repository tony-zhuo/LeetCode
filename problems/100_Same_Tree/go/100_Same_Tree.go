package problems

import datastructures "leet-code/data_structures/binary_tree/go"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *datastructures.TreeNode, q *datastructures.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
