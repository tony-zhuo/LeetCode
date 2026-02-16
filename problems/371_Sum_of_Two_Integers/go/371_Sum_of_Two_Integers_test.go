package problems

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic positive sum - 1 + 2",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "basic positive sum - 2 + 3",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
		{
			name: "one zero - 5 + 0",
			args: args{
				a: 5,
				b: 0,
			},
			want: 5,
		},
		{
			name: "one zero - 0 + 7",
			args: struct {
				a int
				b int
			}{a: 0, b: 7},
			want: 7,
		},
		{
			name: "both zeros",
			args: struct {
				a int
				b int
			}{a: 0, b: 0},
			want: 0,
		},
		{
			name: "negative and positive - (-1) + 1",
			args: struct {
				a int
				b int
			}{a: -1, b: 1},
			want: 0,
		},
		{
			name: "negative and positive - 5 + (-3)",
			args: struct {
				a int
				b int
			}{a: 5, b: -3},
			want: 2,
		},
		{
			name: "positive and negative - (-5) + 3",
			args: struct {
				a int
				b int
			}{a: -5, b: 3},
			want: -2,
		},
		{
			name: "both negative - (-1) + (-2)",
			args: struct {
				a int
				b int
			}{a: -1, b: -2},
			want: -3,
		},
		{
			name: "both negative - (-10) + (-15)",
			args: struct {
				a int
				b int
			}{a: -10, b: -15},
			want: -25,
		},
		{
			name: "larger positive numbers",
			args: struct {
				a int
				b int
			}{a: 100, b: 200},
			want: 300,
		},
		{
			name: "power of 2 numbers",
			args: struct {
				a int
				b int
			}{a: 8, b: 16},
			want: 24,
		},
		{
			name: "large numbers with carry",
			args: struct {
				a int
				b int
			}{a: 999, b: 1},
			want: 1000,
		},
		{
			name: "same numbers",
			args: struct {
				a int
				b int
			}{a: 42, b: 42},
			want: 84,
		},
		{
			name: "bit manipulation edge case",
			args: struct {
				a int
				b int
			}{a: 15, b: 1},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
