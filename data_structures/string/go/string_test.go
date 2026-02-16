package datastructures

import (
	"errors"
	"testing"
)

func TestStringBuilder_AppendAndString(t *testing.T) {
	tests := []struct {
		name    string
		appends []string
		want    string
	}{
		{
			name:    "single append",
			appends: []string{"hello"},
			want:    "hello",
		},
		{
			name:    "multiple appends",
			appends: []string{"hello", " ", "world"},
			want:    "hello world",
		},
		{
			name:    "empty appends",
			appends: []string{"", "a", ""},
			want:    "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilder()
			for _, s := range tt.appends {
				sb.Append(s)
			}
			if got := sb.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStringBuilder_NewStringBuilderFrom(t *testing.T) {
	sb := NewStringBuilderFrom("hello")
	if got := sb.String(); got != "hello" {
		t.Errorf("String() = %q, want %q", got, "hello")
	}
	if got := sb.Len(); got != 5 {
		t.Errorf("Len() = %d, want 5", got)
	}
}

func TestStringBuilder_InsertAt(t *testing.T) {
	tests := []struct {
		name    string
		initial string
		index   int
		insert  string
		want    string
	}{
		{name: "insert at beginning", initial: "world", index: 0, insert: "hello ", want: "hello world"},
		{name: "insert at end", initial: "hello", index: 5, insert: " world", want: "hello world"},
		{name: "insert at middle", initial: "helo", index: 2, insert: "l", want: "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			if err := sb.InsertAt(tt.index, tt.insert); err != nil {
				t.Fatalf("InsertAt() unexpected error: %v", err)
			}
			if got := sb.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStringBuilder_InsertAt_OutOfRange(t *testing.T) {
	sb := NewStringBuilderFrom("abc")
	if err := sb.InsertAt(-1, "x"); !errors.Is(err, ErrStringIndexOutOfRange) {
		t.Errorf("InsertAt(-1) error = %v, want %v", err, ErrStringIndexOutOfRange)
	}
	if err := sb.InsertAt(4, "x"); !errors.Is(err, ErrStringIndexOutOfRange) {
		t.Errorf("InsertAt(4) error = %v, want %v", err, ErrStringIndexOutOfRange)
	}
}

func TestStringBuilder_DeleteRange(t *testing.T) {
	tests := []struct {
		name    string
		initial string
		start   int
		end     int
		want    string
	}{
		{name: "delete from beginning", initial: "hello world", start: 0, end: 6, want: "world"},
		{name: "delete from end", initial: "hello world", start: 5, end: 11, want: "hello"},
		{name: "delete from middle", initial: "hello world", start: 2, end: 5, want: "he world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			if err := sb.DeleteRange(tt.start, tt.end); err != nil {
				t.Fatalf("DeleteRange() unexpected error: %v", err)
			}
			if got := sb.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStringBuilder_CharAt(t *testing.T) {
	sb := NewStringBuilderFrom("abc")
	tests := []struct {
		name  string
		index int
		want  byte
	}{
		{name: "first char", index: 0, want: 'a'},
		{name: "last char", index: 2, want: 'c'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sb.CharAt(tt.index)
			if err != nil {
				t.Fatalf("CharAt(%d) unexpected error: %v", tt.index, err)
			}
			if got != tt.want {
				t.Errorf("CharAt(%d) = %c, want %c", tt.index, got, tt.want)
			}
		})
	}
}

func TestStringBuilder_CharAt_OutOfRange(t *testing.T) {
	sb := NewStringBuilderFrom("abc")
	_, err := sb.CharAt(3)
	if !errors.Is(err, ErrStringIndexOutOfRange) {
		t.Errorf("CharAt(3) error = %v, want %v", err, ErrStringIndexOutOfRange)
	}
}

func TestStringBuilder_Substring(t *testing.T) {
	tests := []struct {
		name    string
		initial string
		start   int
		end     int
		want    string
	}{
		{name: "full string", initial: "hello", start: 0, end: 5, want: "hello"},
		{name: "prefix", initial: "hello", start: 0, end: 3, want: "hel"},
		{name: "suffix", initial: "hello", start: 2, end: 5, want: "llo"},
		{name: "middle", initial: "hello", start: 1, end: 4, want: "ell"},
		{name: "empty substring", initial: "hello", start: 2, end: 2, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			got, err := sb.Substring(tt.start, tt.end)
			if err != nil {
				t.Fatalf("Substring() unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Substring(%d, %d) = %q, want %q", tt.start, tt.end, got, tt.want)
			}
		})
	}
}

func TestStringBuilder_IndexOf(t *testing.T) {
	tests := []struct {
		name    string
		initial string
		target  string
		want    int
	}{
		{name: "found at start", initial: "hello world", target: "hello", want: 0},
		{name: "found at end", initial: "hello world", target: "world", want: 6},
		{name: "found in middle", initial: "hello world", target: "lo w", want: 3},
		{name: "not found", initial: "hello world", target: "xyz", want: -1},
		{name: "single char found", initial: "abc", target: "b", want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			if got := sb.IndexOf(tt.target); got != tt.want {
				t.Errorf("IndexOf(%q) = %d, want %d", tt.target, got, tt.want)
			}
		})
	}
}

func TestStringBuilder_Contains(t *testing.T) {
	sb := NewStringBuilderFrom("hello world")
	if !sb.Contains("world") {
		t.Error("Contains(\"world\") = false, want true")
	}
	if sb.Contains("xyz") {
		t.Error("Contains(\"xyz\") = true, want false")
	}
}

func TestStringBuilder_Replace(t *testing.T) {
	tests := []struct {
		name    string
		initial string
		old     string
		new     string
		want    string
		found   bool
	}{
		{name: "replace found", initial: "hello world", old: "world", new: "go", want: "hello go", found: true},
		{name: "replace not found", initial: "hello world", old: "xyz", new: "go", want: "hello world", found: false},
		{name: "replace first only", initial: "aaa", old: "a", new: "b", want: "baa", found: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			got := sb.Replace(tt.old, tt.new)
			if got != tt.found {
				t.Errorf("Replace() = %v, want %v", got, tt.found)
			}
			if sb.String() != tt.want {
				t.Errorf("after Replace: %q, want %q", sb.String(), tt.want)
			}
		})
	}
}

func TestStringBuilder_ReplaceAll(t *testing.T) {
	tests := []struct {
		name      string
		initial   string
		old       string
		new       string
		want      string
		wantCount int
	}{
		{name: "replace all occurrences", initial: "aabaa", old: "a", new: "x", want: "xxbxx", wantCount: 4},
		{name: "no occurrences", initial: "hello", old: "z", new: "x", want: "hello", wantCount: 0},
		{name: "replace with longer", initial: "abab", old: "ab", new: "xyz", want: "xyzxyz", wantCount: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			count := sb.ReplaceAll(tt.old, tt.new)
			if count != tt.wantCount {
				t.Errorf("ReplaceAll() count = %d, want %d", count, tt.wantCount)
			}
			if sb.String() != tt.want {
				t.Errorf("after ReplaceAll: %q, want %q", sb.String(), tt.want)
			}
		})
	}
}

func TestStringBuilder_Reverse(t *testing.T) {
	tests := []struct {
		name    string
		initial string
		want    string
	}{
		{name: "odd length", initial: "hello", want: "olleh"},
		{name: "even length", initial: "abcd", want: "dcba"},
		{name: "single char", initial: "a", want: "a"},
		{name: "empty", initial: "", want: ""},
		{name: "palindrome", initial: "racecar", want: "racecar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			sb.Reverse()
			if got := sb.String(); got != tt.want {
				t.Errorf("after Reverse: %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStringBuilder_IsPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		initial string
		want    bool
	}{
		{name: "palindrome odd", initial: "racecar", want: true},
		{name: "palindrome even", initial: "abba", want: true},
		{name: "not palindrome", initial: "hello", want: false},
		{name: "single char", initial: "a", want: true},
		{name: "empty string", initial: "", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			if got := sb.IsPalindrome(); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringBuilder_CountChar(t *testing.T) {
	sb := NewStringBuilderFrom("hello world")
	if got := sb.CountChar('l'); got != 3 {
		t.Errorf("CountChar('l') = %d, want 3", got)
	}
	if got := sb.CountChar('z'); got != 0 {
		t.Errorf("CountChar('z') = %d, want 0", got)
	}
}

func TestStringBuilder_ToCharArray(t *testing.T) {
	sb := NewStringBuilderFrom("abc")
	arr := sb.ToCharArray()
	want := []byte{'a', 'b', 'c'}
	for i, b := range arr {
		if b != want[i] {
			t.Errorf("ToCharArray()[%d] = %c, want %c", i, b, want[i])
		}
	}
	// Verify it's a copy
	arr[0] = 'z'
	if ch, _ := sb.CharAt(0); ch != 'a' {
		t.Error("ToCharArray should return a copy, not a reference")
	}
}

func TestStringBuilder_Len(t *testing.T) {
	tests := []struct {
		name    string
		initial string
		want    int
	}{
		{name: "empty", initial: "", want: 0},
		{name: "non-empty", initial: "hello", want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFrom(tt.initial)
			if got := sb.Len(); got != tt.want {
				t.Errorf("Len() = %d, want %d", got, tt.want)
			}
		})
	}
}
