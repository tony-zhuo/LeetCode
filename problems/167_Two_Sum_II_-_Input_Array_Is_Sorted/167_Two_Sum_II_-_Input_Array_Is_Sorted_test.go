package problems

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1: Two numbers sum to 9",
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			name: "Example 2: Two numbers sum to 6",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			name: "Example 3: Negative number",
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
		{
			name: "Two elements",
			args: args{
				numbers: []int{1, 2},
				target:  3,
			},
			want: []int{1, 2},
		},
		{
			name: "Target at the end",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
				target:  9,
			},
			want: []int{4, 5},
		},
		{
			name: "Negative numbers",
			args: args{
				numbers: []int{-5, -3, -1, 2, 4},
				target:  -8,
			},
			want: []int{1, 2},
		},
		{
			name: "Large array with answer in the middle",
			args: args{
				numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				target:  11,
			},
			want: []int{1, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}