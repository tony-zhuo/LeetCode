package problems

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1: Multiple words with single space",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "Example 2: Leading and trailing spaces",
			args: args{
				s: "  hello world  ",
			},
			want: "world hello",
		},
		{
			name: "Example 3: Multiple spaces between words",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
		{
			name: "Single word",
			args: args{
				s: "hello",
			},
			want: "hello",
		},
		{
			name: "Single word with leading spaces",
			args: args{
				s: "   hello",
			},
			want: "hello",
		},
		{
			name: "Single word with trailing spaces",
			args: args{
				s: "hello   ",
			},
			want: "hello",
		},
		{
			name: "Single word with leading and trailing spaces",
			args: args{
				s: "  hello  ",
			},
			want: "hello",
		},
		{
			name: "Two words",
			args: args{
				s: "hello world",
			},
			want: "world hello",
		},
		{
			name: "Multiple spaces everywhere",
			args: args{
				s: "  Bob    Loves  Alice   ",
			},
			want: "Alice Loves Bob",
		},
		{
			name: "Single character word",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "Multiple single character words",
			args: args{
				s: "a b c",
			},
			want: "c b a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
