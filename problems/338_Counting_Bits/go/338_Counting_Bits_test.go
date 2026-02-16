package problems

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "n = 0",
			args: args{n: 0},
			want: []int{0}, // [0]
		},
		{
			name: "n = 1",
			args: args{n: 1},
			want: []int{0, 1}, // [0b0, 0b1]
		},
		{
			name: "n = 2",
			args: args{n: 2},
			want: []int{0, 1, 1}, // [0b0, 0b1, 0b10]
		},
		{
			name: "n = 3",
			args: args{n: 3},
			want: []int{0, 1, 1, 2}, // [0b0, 0b1, 0b10, 0b11]
		},
		{
			name: "n = 4",
			args: args{n: 4},
			want: []int{0, 1, 1, 2, 1}, // [0b0, 0b1, 0b10, 0b11, 0b100]
		},
		{
			name: "n = 5",
			args: args{n: 5},
			want: []int{0, 1, 1, 2, 1, 2}, // [0b0, 0b1, 0b10, 0b11, 0b100, 0b101]
		},
		{
			name: "n = 7 (power of 2 minus 1)",
			args: args{n: 7},
			want: []int{0, 1, 1, 2, 1, 2, 2, 3}, // [0b0, 0b1, 0b10, 0b11, 0b100, 0b101, 0b110, 0b111]
		},
		{
			name: "n = 8 (power of 2)",
			args: args{n: 8},
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1}, // [..., 0b1000]
		},
		{
			name: "n = 15 (all bits set in 4-bit)",
			args: args{n: 15},
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}, // 0 to 15
		},
		{
			name: "n = 16 (power of 2)",
			args: args{n: 16},
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1}, // [..., 0b10000]
		},
		{
			name: "n = 31 (large case)",
			args: args{n: 31},
			want: []int{
				0, // 0b0
				1, // 0b1
				1, // 0b10
				2, // 0b11
				1, // 0b100
				2, // 0b101
				2, // 0b110
				3, // 0b111
				1, // 0b1000
				2, // 0b1001
				2, // 0b1010
				3, // 0b1011
				2, // 0b1100
				3, // 0b1101
				3, // 0b1110
				4, // 0b1111
				1, // 0b10000
				2, // 0b10001
				2, // 0b10010
				3, // 0b10011
				2, // 0b10100
				3, // 0b10101
				3, // 0b10110
				4, // 0b10111
				2, // 0b11000
				3, // 0b11001
				3, // 0b11010
				4, // 0b11011
				3, // 0b11100
				4, // 0b11101
				4, // 0b11110
				5, // 0b11111
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
