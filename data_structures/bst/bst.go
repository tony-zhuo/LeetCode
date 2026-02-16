package datastructures

import "errors"

var ErrEmptyTree = errors.New("tree is empty")
var ErrNotFound = errors.New("value not found")

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root *BSTNode
	size int
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) Insert(val int) {
	b.Root = insert(b.Root, val, &b.size)
}

func insert(node *BSTNode, val int, size *int) *BSTNode {
	if node == nil {
		*size++
		return &BSTNode{Val: val}
	}

	if val < node.Val {
		node.Left = insert(node.Left, val, size)
	} else if val > node.Val {
		node.Right = insert(node.Right, val, size)
	}
	// duplicate: do nothing

	return node
}

func (b *BST) Delete(val int) bool {
	var deleted bool
	b.Root = deleteNode(b.Root, val, &deleted)
	if deleted {
		b.size--
	}
	return deleted
}

func deleteNode(node *BSTNode, val int, deleted *bool) *BSTNode {
	if node == nil {
		return nil
	}

	if val < node.Val {
		node.Left = deleteNode(node.Left, val, deleted)
	} else if val > node.Val {
		node.Right = deleteNode(node.Right, val, deleted)
	} else {
		*deleted = true

		// Case 1: leaf node
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// Case 2: one child
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// Case 3: two children - use in-order successor
		successor := findMin(node.Right)
		node.Val = successor.Val
		dummy := false
		node.Right = deleteNode(node.Right, successor.Val, &dummy)
	}

	return node
}

func findMin(node *BSTNode) *BSTNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func findMax(node *BSTNode) *BSTNode {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

func (b *BST) Search(val int) bool {
	return search(b.Root, val)
}

func search(node *BSTNode, val int) bool {
	if node == nil {
		return false
	}

	if val < node.Val {
		return search(node.Left, val)
	} else if val > node.Val {
		return search(node.Right, val)
	}

	return true
}

func (b *BST) Min() (int, error) {
	if b.Root == nil {
		return 0, ErrEmptyTree
	}
	return findMin(b.Root).Val, nil
}

func (b *BST) Max() (int, error) {
	if b.Root == nil {
		return 0, ErrEmptyTree
	}
	return findMax(b.Root).Val, nil
}

func (b *BST) InorderTraversal() []int {
	result := []int{}
	inorder(b.Root, &result)
	return result
}

func inorder(node *BSTNode, result *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, result)
	*result = append(*result, node.Val)
	inorder(node.Right, result)
}

func (b *BST) Size() int {
	return b.size
}
