package problems

import (
	datastructures "leet-code/data_structures/linked_list/go"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *datastructures.ListNode
	}
	tests := []struct {
		name string
		args args
		want *datastructures.ListNode
	}{
		{
			name: "",
			args: args{
				head: &datastructures.ListNode{
					Val: 1,
					Next: &datastructures.ListNode{
						Val: 2,
						Next: &datastructures.ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			want: &datastructures.ListNode{
				Val: 3,
				Next: &datastructures.ListNode{
					Val: 2,
					Next: &datastructures.ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
