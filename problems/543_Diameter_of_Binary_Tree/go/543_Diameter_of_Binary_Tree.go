package problems

import datastructures "leet-code/data_structures/binary_tree/go"

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func diameterOfBinaryTree(root *datastructures.TreeNode) int {
	ans := 0

	var depth func(node *datastructures.TreeNode) int
	depth = func(node *datastructures.TreeNode) int {
		if node == nil {
			return 0
		}

		left := depth(node.Left)
		right := depth(node.Right)
		ans = max(ans, left+right)
		return max(left, right) + 1
	}

	depth(root)

	return ans
}
