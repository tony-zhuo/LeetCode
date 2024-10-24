package problems

import (
	"leet-code/structure"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "",
			args: args{
				head: &structure.ListNode{
					Val: 1,
					Next: &structure.ListNode{
						Val: 2,
						Next: &structure.ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			want: &structure.ListNode{
				Val: 3,
				Next: &structure.ListNode{
					Val: 2,
					Next: &structure.ListNode{
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
