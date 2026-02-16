package problems

import datastructures "leet-code/data_structures/binary_tree/go"

func buildTree(preorder []int, inorder []int) *datastructures.TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}

	mid := -1
	root := &datastructures.TreeNode{Val: preorder[0]}
	for i, v := range inorder {
		if root.Val == v {
			mid = i
			break
		}
	}
	if mid == -1 {
		return nil
	}

	root.Left = buildTree(preorder[1:], inorder[:mid])
	root.Right = buildTree(preorder[1+mid:], inorder[mid+1:])

	return root
}
