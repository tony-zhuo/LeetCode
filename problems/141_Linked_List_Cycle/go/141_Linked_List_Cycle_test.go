package problems

import (
	datastructures "leet-code/data_structures/linked_list/go"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *datastructures.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "cycle at tail pointing to second node",
			args: args{
				head: func() *datastructures.ListNode {
					// [3,2,0,-4] with tail connecting to node index 1
					node1 := &datastructures.ListNode{Val: 3}
					node2 := &datastructures.ListNode{Val: 2}
					node3 := &datastructures.ListNode{Val: 0}
					node4 := &datastructures.ListNode{Val: -4}
					
					node1.Next = node2
					node2.Next = node3
					node3.Next = node4
					node4.Next = node2 // cycle: tail points to node2
					
					return node1
				}(),
			},
			want: true,
		},
		{
			name: "cycle at tail pointing to head",
			args: args{
				head: func() *datastructures.ListNode {
					// [1,2] with tail connecting to node index 0
					node1 := &datastructures.ListNode{Val: 1}
					node2 := &datastructures.ListNode{Val: 2}
					
					node1.Next = node2
					node2.Next = node1 // cycle: tail points to head
					
					return node1
				}(),
			},
			want: true,
		},
		{
			name: "no cycle - single node",
			args: args{
				head: &datastructures.ListNode{Val: 1, Next: nil},
			},
			want: false,
		},
		{
			name: "no cycle - multiple nodes",
			args: args{
				head: func() *datastructures.ListNode {
					// [1,2,3,4] with no cycle
					node1 := &datastructures.ListNode{Val: 1}
					node2 := &datastructures.ListNode{Val: 2}
					node3 := &datastructures.ListNode{Val: 3}
					node4 := &datastructures.ListNode{Val: 4}
					
					node1.Next = node2
					node2.Next = node3
					node3.Next = node4
					// node4.Next = nil (no cycle)
					
					return node1
				}(),
			},
			want: false,
		},
		{
			name: "empty list",
			args: args{
				head: nil,
			},
			want: false,
		},
		{
			name: "self-cycle",
			args: args{
				head: func() *datastructures.ListNode {
					// Single node pointing to itself
					node := &datastructures.ListNode{Val: 1}
					node.Next = node
					return node
				}(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
