package structure

type ListNode struct {
	Val  int
	Next *ListNode
}

func Arr2Node(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  input[0],
		Next: nil,
	}
	curr := head

	for i := 1; i < len(input); i++ {
		curr.Next = &ListNode{
			Val:  input[i],
			Next: nil,
		}

		curr = curr.Next
	}

	return head
}
