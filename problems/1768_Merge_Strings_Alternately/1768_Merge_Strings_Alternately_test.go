package problems

import "testing"

func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1: Equal length strings",
			args: args{
				word1: "abc",
				word2: "pqr",
			},
			want: "apbqcr",
		},
		{
			name: "Example 2: word1 longer",
			args: args{
				word1: "ab",
				word2: "pqrs",
			},
			want: "apbqrs",
		},
		{
			name: "Example 3: word2 longer",
			args: args{
				word1: "abcd",
				word2: "pq",
			},
			want: "apbqcd",
		},
		{
			name: "Empty word1",
			args: args{
				word1: "",
				word2: "abc",
			},
			want: "abc",
		},
		{
			name: "Empty word2",
			args: args{
				word1: "abc",
				word2: "",
			},
			want: "abc",
		},
		{
			name: "Both empty",
			args: args{
				word1: "",
				word2: "",
			},
			want: "",
		},
		{
			name: "Single characters",
			args: args{
				word1: "a",
				word2: "b",
			},
			want: "ab",
		},
		{
			name: "Long strings with different lengths",
			args: args{
				word1: "hello",
				word2: "worldxyz",
			},
			want: "hweolrllodxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
