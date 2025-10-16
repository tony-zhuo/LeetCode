package problems

import (
	"leet-code/structure"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "cycle at tail pointing to second node",
			args: args{
				head: func() *structure.ListNode {
					// [3,2,0,-4] with tail connecting to node index 1
					node1 := &structure.ListNode{Val: 3}
					node2 := &structure.ListNode{Val: 2}
					node3 := &structure.ListNode{Val: 0}
					node4 := &structure.ListNode{Val: -4}
					
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
				head: func() *structure.ListNode {
					// [1,2] with tail connecting to node index 0
					node1 := &structure.ListNode{Val: 1}
					node2 := &structure.ListNode{Val: 2}
					
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
				head: &structure.ListNode{Val: 1, Next: nil},
			},
			want: false,
		},
		{
			name: "no cycle - multiple nodes",
			args: args{
				head: func() *structure.ListNode {
					// [1,2,3,4] with no cycle
					node1 := &structure.ListNode{Val: 1}
					node2 := &structure.ListNode{Val: 2}
					node3 := &structure.ListNode{Val: 3}
					node4 := &structure.ListNode{Val: 4}
					
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
				head: func() *structure.ListNode {
					// Single node pointing to itself
					node := &structure.ListNode{Val: 1}
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
