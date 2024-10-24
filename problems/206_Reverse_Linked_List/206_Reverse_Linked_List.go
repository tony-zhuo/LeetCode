package problems

import "leet-code/structure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *structure.ListNode) *structure.ListNode {
	if head == nil {
		return nil
	}
	curr := head

	var previous *structure.ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = previous
		previous = curr
		curr = next
	}

	return previous
}
