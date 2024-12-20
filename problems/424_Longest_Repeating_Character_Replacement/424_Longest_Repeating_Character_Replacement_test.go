package problems

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "AAAB",
				k: 0,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "BAAAB",
				k: 2,
			},
			want: 5,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
