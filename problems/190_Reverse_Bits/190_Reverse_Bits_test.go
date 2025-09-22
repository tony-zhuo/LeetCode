package problems

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{n: 43261596},  // 00000010100101000001111010011100
			want: 964176192,         // 00111001011110000010100101000000
		},
		{
			name: "Example 2", 
			args: args{n: 4294967293}, // 11111111111111111111111111111101
			want: 3221225471,          // 10111111111111111111111111111111
		},
		{
			name: "All zeros",
			args: args{n: 0},          // 00000000000000000000000000000000
			want: 0,                   // 00000000000000000000000000000000
		},
		{
			name: "All ones",
			args: args{n: 4294967295}, // 11111111111111111111111111111111
			want: 4294967295,          // 11111111111111111111111111111111
		},
		{
			name: "Single bit at start",
			args: args{n: 1},          // 00000000000000000000000000000001
			want: 2147483648,          // 10000000000000000000000000000000
		},
		{
			name: "Single bit at end",
			args: args{n: 2147483648}, // 10000000000000000000000000000000
			want: 1,                   // 00000000000000000000000000000001
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.n); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
