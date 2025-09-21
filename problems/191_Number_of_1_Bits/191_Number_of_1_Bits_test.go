package problems

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single bit set",
			args: args{n: 1}, // 0b1
			want: 1,
		},
		{
			name: "multiple bits set",
			args: args{n: 11}, // 0b1011
			want: 3,
		},
		{
			name: "all bits set in byte",
			args: args{n: 255}, // 0b11111111
			want: 8,
		},
		{
			name: "power of 2",
			args: args{n: 128}, // 0b10000000
			want: 1,
		},
		{
			name: "zero",
			args: args{n: 0}, // 0b0
			want: 0,
		},
		{
			name: "large number with many bits",
			args: args{n: 1023}, // 0b1111111111 (10 bits)
			want: 10,
		},
		{
			name: "alternating bits pattern",
			args: args{n: 85}, // 0b1010101
			want: 4,
		},
		{
			name: "another alternating pattern",
			args: args{n: 170}, // 0b10101010
			want: 4,
		},
		{
			name: "power of 2 minus 1",
			args: args{n: 31}, // 0b11111
			want: 5,
		},
		{
			name: "large power of 2",
			args: args{n: 1024}, // 0b10000000000
			want: 1,
		},
		{
			name: "maximum 16-bit value",
			args: args{n: 65535}, // 0b1111111111111111
			want: 16,
		},
		{
			name: "sparse bits",
			args: args{n: 4369}, // 0b1000100010001
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.n); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
