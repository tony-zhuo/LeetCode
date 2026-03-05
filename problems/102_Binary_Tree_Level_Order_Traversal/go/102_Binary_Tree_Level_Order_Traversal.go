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

func levelOrder_BFS(root *datastructures.TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*datastructures.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		res = append(res, level)
	}

	return res
}
