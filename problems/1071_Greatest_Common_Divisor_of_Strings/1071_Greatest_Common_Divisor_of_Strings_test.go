package problems

import "testing"

func Test_gcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{
			name: "Example 1",
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			name: "Example 2",
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			name: "Example 3",
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
		{
			name: "Same strings",
			str1: "ABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			name: "No common divisor",
			str1: "ABCDEF",
			str2: "GHIJKL",
			want: "",
		},
		{
			name: "One character GCD",
			str1: "AAA",
			str2: "AA",
			want: "A",
		},
		{
			name: "Empty result case",
			str1: "ABAB",
			str2: "BABA",
			want: "",
		},
		{
			name: "Single character",
			str1: "A",
			str2: "A",
			want: "A",
		},
		{
			name: "Different single characters",
			str1: "A",
			str2: "B",
			want: "",
		},
		{
			name: "Complex pattern",
			str1: "ABCDEFABCDEF",
			str2: "ABCDEF",
			want: "ABCDEF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.str1, tt.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
