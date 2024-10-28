package structure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Slice2BinaryTree(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil // Return nil if the slice is empty
	}

	// Create the root node
	root := &TreeNode{Val: *arr[0]}
	queue := []*TreeNode{root} // Queue to help build the tree level by level
	index := 1                 // Index to track the current element in the slice

	// Use the queue to build the binary tree
	for index < len(arr) {
		// Pop the front node from the queue
		current := queue[0]
		queue = queue[1:]

		// Insert the left child
		if index < len(arr) && arr[index] != nil {
			current.Left = &TreeNode{Val: *arr[index]}
			queue = append(queue, current.Left)
		}
		index++

		// Insert the right child
		if index < len(arr) && arr[index] != nil {
			current.Right = &TreeNode{Val: *arr[index]}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}
