package problems

import datastructures "leet-code/data_structures/binary_tree/go"

func levelOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	m := make(map[int][]int)
	scanTree(root, 0, m)

	res := make([][]int, len(m))
	for k, v := range m {
		res[k] = v
	}

	return res
}

func scanTree(root *datastructures.TreeNode, level int, m map[int][]int) {
	if root == nil {
		return
	}
	m[level] = append(m[level], root.Val)
	scanTree(root.Left, level+1, m)
	scanTree(root.Right, level+1, m)

	return
}
