package structure

import (
	"reflect"
	"testing"
)

func TestArr2Node(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				input: []int{1, 2, 3},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Arr2Node(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Arr2Node() = %v, want %v", got, tt.want)
			}
		})
	}
}
