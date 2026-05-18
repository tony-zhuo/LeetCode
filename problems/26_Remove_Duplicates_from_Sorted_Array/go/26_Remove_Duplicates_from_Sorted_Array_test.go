package problems

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		wantK    int
		wantNums []int
	}{
		{
			name:     "Example 1: [1,1,2]",
			args:     args{nums: []int{1, 1, 2}},
			wantK:    2,
			wantNums: []int{1, 2},
		},
		{
			name:     "Example 2: [0,0,1,1,1,2,2,3,3,4]",
			args:     args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			wantK:    5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "Single element",
			args:     args{nums: []int{1}},
			wantK:    1,
			wantNums: []int{1},
		},
		{
			name:     "All duplicates",
			args:     args{nums: []int{2, 2, 2, 2, 2}},
			wantK:    1,
			wantNums: []int{2},
		},
		{
			name:     "No duplicates",
			args:     args{nums: []int{1, 2, 3, 4, 5}},
			wantK:    5,
			wantNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Negative numbers",
			args:     args{nums: []int{-3, -3, -2, -1, -1, 0}},
			wantK:    4,
			wantNums: []int{-3, -2, -1, 0},
		},
		{
			name:     "Two elements same",
			args:     args{nums: []int{1, 1}},
			wantK:    1,
			wantNums: []int{1},
		},
		{
			name:     "Two elements different",
			args:     args{nums: []int{1, 2}},
			wantK:    2,
			wantNums: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.wantK {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.wantK)
			}
			if !reflect.DeepEqual(tt.args.nums[:got], tt.wantNums) {
				t.Errorf("nums[:k] = %v, want %v", tt.args.nums[:got], tt.wantNums)
			}
		})
	}
}
