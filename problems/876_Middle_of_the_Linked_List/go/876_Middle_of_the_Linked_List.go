package problems

import datastructures "leet-code/data_structures/linked_list/go"

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func middleNode_two_pass(head *datastructures.ListNode) *datastructures.ListNode {
	cnt := 0
	node := head
	for node != nil {
		node = node.Next
		cnt++
	}

	mid := cnt / 2
	cnt = 0
	for cnt < mid {
		head = head.Next
		cnt++
	}

	return head
}

func middleNode_slow_fast(head *datastructures.ListNode) *datastructures.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
