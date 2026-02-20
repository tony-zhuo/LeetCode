package problems

import "testing"

var cases = []struct {
	name string
	a    string
	b    string
	want string
}{
	{name: "example 1", a: "11", b: "1", want: "100"},
	{name: "example 2", a: "1010", b: "1011", want: "10101"},
	{name: "both zero", a: "0", b: "0", want: "0"},
	{name: "carry chain", a: "1111", b: "1", want: "10000"},
	{name: "different lengths", a: "1", b: "111", want: "1000"},
}

func Test_addBinary(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := addBinary(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("addBinary(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
