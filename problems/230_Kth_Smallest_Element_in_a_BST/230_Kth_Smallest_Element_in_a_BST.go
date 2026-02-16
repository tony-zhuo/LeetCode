package problems

import (
	datastructures "leet-code/data_structures/binary_tree"
	"sort"
)

func kthSmallest(root *datastructures.TreeNode, k int) int {
	tmp := orderTree(root)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	return tmp[k-1]
}

func orderTree(root *datastructures.TreeNode) []int {
	tmp := make([]int, 0)
	if root == nil {
		return tmp
	}

	tmp = append(tmp, root.Val)

	leftTmp := orderTree(root.Left)
	tmp = append(tmp, leftTmp...)

	rightTmp := orderTree(root.Right)
	tmp = append(tmp, rightTmp...)

	return tmp
}
