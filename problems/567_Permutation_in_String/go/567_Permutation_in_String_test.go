package problems

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1: s2 contains permutation 'ba'",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
		{
			name: "Example 2: no permutation found",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},
		{
			name: "Exact match",
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: true,
		},
		{
			name: "s1 longer than s2",
			args: args{
				s1: "abcd",
				s2: "ab",
			},
			want: false,
		},
		{
			name: "Single character match",
			args: args{
				s1: "a",
				s2: "a",
			},
			want: true,
		},
		{
			name: "Single character no match",
			args: args{
				s1: "a",
				s2: "b",
			},
			want: false,
		},
		{
			name: "Permutation at the end",
			args: args{
				s1: "adc",
				s2: "dcda",
			},
			want: true,
		},
		{
			name: "Repeated characters",
			args: args{
				s1: "aab",
				s2: "cabaab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
