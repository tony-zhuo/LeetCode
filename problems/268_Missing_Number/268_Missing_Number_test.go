package problems

import "testing"

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "example 1 - missing 2",
			args: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "example 2 - missing 2",
			args: []int{0, 1},
			want: 2,
		},
		{
			name: "example 3 - missing 8",
			args: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
		{
			name: "single element - missing 0",
			args: []int{1},
			want: 0,
		},
		{
			name: "single element - missing 1",
			args: []int{0},
			want: 1,
		},
		{
			name: "missing first number",
			args: []int{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			name: "missing last number",
			args: []int{0, 1, 2, 3, 4},
			want: 5,
		},
		{
			name: "missing middle number",
			args: []int{0, 1, 3, 4, 5},
			want: 2,
		},
		{
			name: "two elements - missing 1",
			args: []int{0, 2},
			want: 1,
		},
		{
			name: "larger array missing middle",
			args: []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
