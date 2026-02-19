package problems

import "testing"

var cases = []struct {
	name string
	n    int
	bad  int
}{
	{name: "example 1", n: 5, bad: 4},
	{name: "example 2", n: 1, bad: 1},
	{name: "first is bad", n: 10, bad: 1},
	{name: "last is bad", n: 10, bad: 10},
	{name: "mid point", n: 100, bad: 50},
	{name: "two versions", n: 2, bad: 2},
}

func Test_firstBadVersion(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			isBadVersion = func(version int) bool {
				return version >= tt.bad
			}
			got := firstBadVersion(tt.n)
			if got != tt.bad {
				t.Errorf("firstBadVersion(%d) = %v, want %v", tt.n, got, tt.bad)
			}
		})
	}
}
