package datastructures

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode creates a new TreeNode with the given value.
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// InorderTraversal returns the inorder traversal of the binary tree (Left, Root, Right).
func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorder(root, &result)
	return result
}

func inorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, result)
	*result = append(*result, node.Val)
	inorder(node.Right, result)
}

// PreorderTraversal returns the preorder traversal of the binary tree (Root, Left, Right).
func PreorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorder(root, &result)
	return result
}

func preorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preorder(node.Left, result)
	preorder(node.Right, result)
}

// PostorderTraversal returns the postorder traversal of the binary tree (Left, Right, Root).
func PostorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorder(root, &result)
	return result
}

func postorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, result)
	postorder(node.Right, result)
	*result = append(*result, node.Val)
}

// LevelOrderTraversal returns the level order traversal of the binary tree.
// Each inner slice contains the values at that level from left to right.
func LevelOrderTraversal(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
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
		result = append(result, level)
	}

	return result
}

// Height returns the height of the binary tree.
// An empty tree has height 0, a single node has height 1.
func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := Height(root.Left)
	rightHeight := Height(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// CountNodes returns the total number of nodes in the binary tree.
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}

// Slice2BinaryTree constructs a binary tree from a level-order slice.
// nil entries represent absent nodes.
func Slice2BinaryTree(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	queue := []*TreeNode{root}
	index := 1

	for index < len(arr) {
		current := queue[0]
		queue = queue[1:]

		if index < len(arr) && arr[index] != nil {
			current.Left = &TreeNode{Val: *arr[index]}
			queue = append(queue, current.Left)
		}
		index++

		if index < len(arr) && arr[index] != nil {
			current.Right = &TreeNode{Val: *arr[index]}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}
