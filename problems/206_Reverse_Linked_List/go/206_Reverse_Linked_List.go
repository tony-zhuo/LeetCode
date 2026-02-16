package problems

import datastructures "leet-code/data_structures/linked_list/go"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *datastructures.ListNode) *datastructures.ListNode {
	if head == nil {
		return nil
	}
	curr := head

	var previous *datastructures.ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = previous
		previous = curr
		curr = next
	}

	return previous
}
