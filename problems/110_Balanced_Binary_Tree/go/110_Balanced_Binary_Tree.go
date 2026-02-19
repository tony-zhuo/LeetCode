package problems

import datastructures "leet-code/data_structures/binary_tree/go"

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *datastructures.TreeNode) bool {
	return hight(root) != -1
}

func hight(node *datastructures.TreeNode) int {
	if node == nil {
		return 0
	}
	left := hight(node.Left)
	right := hight(node.Right)

	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
