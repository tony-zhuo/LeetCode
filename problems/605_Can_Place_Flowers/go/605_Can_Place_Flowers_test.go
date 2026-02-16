package problems

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1 - can place flowers",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Example 2 - cannot place flowers",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "Single empty spot",
			args: args{
				flowerbed: []int{0},
				n:         1,
			},
			want: true,
		},
		{
			name: "Single occupied spot",
			args: args{
				flowerbed: []int{1},
				n:         1,
			},
			want: false,
		},
		{
			name: "All empty - can place all",
			args: args{
				flowerbed: []int{0, 0, 0, 0, 0},
				n:         3,
			},
			want: true,
		},
		{
			name: "All empty - cannot place too many",
			args: args{
				flowerbed: []int{0, 0, 0, 0, 0},
				n:         4,
			},
			want: false,
		},
		{
			name: "All occupied",
			args: args{
				flowerbed: []int{1, 1, 1, 1, 1},
				n:         1,
			},
			want: false,
		},
		{
			name: "No flowers to place",
			args: args{
				flowerbed: []int{1, 0, 1, 0, 1},
				n:         0,
			},
			want: true,
		},
		{
			name: "Start with empty spots",
			args: args{
				flowerbed: []int{0, 0, 1, 0, 0},
				n:         1,
			},
			want: true,
		},
		{
			name: "End with empty spots",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 0},
				n:         2,
			},
			want: true,
		},
		{
			name: "Alternating pattern",
			args: args{
				flowerbed: []int{1, 0, 1, 0, 1, 0, 1},
				n:         1,
			},
			want: false,
		},
		{
			name: "Two consecutive empty at start",
			args: args{
				flowerbed: []int{0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Two consecutive empty at end",
			args: args{
				flowerbed: []int{1, 0, 0},
				n:         1,
			},
			want: true,
		},
		{
			name: "Large array with many opportunities",
			args: args{
				flowerbed: []int{0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
				n:         3,
			},
			want: true,
		},
		{
			name: "Large array insufficient space",
			args: args{
				flowerbed: []int{0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
				n:         4,
			},
			want: false,
		},
		{
			name: "Edge case - two elements both empty",
			args: args{
				flowerbed: []int{0, 0},
				n:         1,
			},
			want: true,
		},
		{
			name: "Edge case - two elements both empty, want two",
			args: args{
				flowerbed: []int{0, 0},
				n:         2,
			},
			want: false,
		},
		{
			name: "Complex pattern",
			args: args{
				flowerbed: []int{0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
				n:         2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
