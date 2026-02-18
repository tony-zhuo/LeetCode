package problems

import datastructures "leet-code/data_structures/binary_tree/go"

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func lowestCommonAncestor(root, p, q *datastructures.TreeNode) *datastructures.TreeNode {
	node := root
	for node != nil {
		if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		} else if p.Val > node.Val && q.Val > node.Val {
			node = node.Right
		} else {
			return node
		}
	}

	return nil
}
