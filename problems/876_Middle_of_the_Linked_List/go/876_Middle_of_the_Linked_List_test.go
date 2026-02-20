package problems

import (
	"testing"

	datastructures "leet-code/data_structures/linked_list/go"
)

func node2arr(head *datastructures.ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

var cases = []struct {
	name string
	head []int
	want []int
}{
	{name: "odd length", head: []int{1, 2, 3, 4, 5}, want: []int{3, 4, 5}},
	{name: "even length", head: []int{1, 2, 3, 4, 5, 6}, want: []int{4, 5, 6}},
	{name: "single node", head: []int{1}, want: []int{1}},
	{name: "two nodes", head: []int{1, 2}, want: []int{2}},
}

func Test_middleNode_two_pass(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.Arr2Node(tt.head)
			got := node2arr(middleNode_two_pass(head))
			if len(got) != len(tt.want) {
				t.Errorf("middleNode_two_pass() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("middleNode_two_pass() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func Test_middleNode_slow_fast(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.Arr2Node(tt.head)
			got := node2arr(middleNode_slow_fast(head))
			if len(got) != len(tt.want) {
				t.Errorf("middleNode_slow_fast() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("middleNode_slow_fast() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
