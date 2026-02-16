package problems

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "simple parens", s: "()", want: true},
		{name: "all types", s: "()[]{}", want: true},
		{name: "mismatch", s: "(]", want: false},
		{name: "interleaved", s: "([)]", want: false},
		{name: "nested", s: "{[]}", want: true},
		{name: "empty", s: "", want: true},
		{name: "unmatched open", s: "(", want: false},
		{name: "unmatched close", s: ")", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
