package problems

import "testing"

var cases = []struct {
	name string
	s    string
	t    string
	want string
}{
	{name: "all different possible", s: "1100", t: "1010", want: "1111"},
	{name: "all zeros", s: "0000", t: "0000", want: "0000"},
	{name: "all ones cancel", s: "1111", t: "1111", want: "0000"},
	{name: "simple pair", s: "10", t: "10", want: "11"},
	{name: "not enough zeros", s: "110", t: "110", want: "101"},
	{name: "single bit", s: "1", t: "0", want: "1"},
}

func Test_maximumXor(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumXor(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("maximumXor(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
