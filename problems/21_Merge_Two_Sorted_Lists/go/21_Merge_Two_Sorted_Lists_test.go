package problems

import (
	"reflect"
	"testing"

	datastructures "leet-code/data_structures/linked_list/go"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{name: "both non-empty", list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, want: []int{1, 1, 2, 3, 4, 4}},
		{name: "both empty", list1: []int{}, list2: []int{}, want: []int{}},
		{name: "first empty", list1: []int{}, list2: []int{0}, want: []int{0}},
		{name: "second empty", list1: []int{1, 2, 3}, list2: []int{}, want: []int{1, 2, 3}},
		{name: "single elements", list1: []int{1}, list2: []int{2}, want: []int{1, 2}},
		{name: "interleaved", list1: []int{1, 3, 5}, list2: []int{2, 4, 6}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "all same values", list1: []int{1, 1, 1}, list2: []int{1, 1}, want: []int{1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := datastructures.Arr2Node(tt.list1)
			l2 := datastructures.Arr2Node(tt.list2)
			got := mergeTwoLists(l1, l2)
			gotArr := node2Arr(got)
			if !reflect.DeepEqual(gotArr, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", gotArr, tt.want)
			}
		})
	}
}

func node2Arr(head *datastructures.ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
