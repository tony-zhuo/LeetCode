package datastructures

import "errors"

var ErrStringIndexOutOfRange = errors.New("string index out of range")

// StringBuilder provides mutable string operations built on a byte slice.
type StringBuilder struct {
	data []byte
}

// NewStringBuilder creates an empty StringBuilder.
func NewStringBuilder() *StringBuilder {
	return &StringBuilder{data: []byte{}}
}

// NewStringBuilderFrom creates a StringBuilder initialized with the given string.
func NewStringBuilderFrom(s string) *StringBuilder {
	data := make([]byte, len(s))
	copy(data, s)
	return &StringBuilder{data: data}
}

// Append adds a string to the end.
func (sb *StringBuilder) Append(s string) {
	sb.data = append(sb.data, s...)
}

// InsertAt inserts a string at the given byte index.
func (sb *StringBuilder) InsertAt(index int, s string) error {
	if index < 0 || index > len(sb.data) {
		return ErrStringIndexOutOfRange
	}
	sb.data = append(sb.data[:index], append([]byte(s), sb.data[index:]...)...)
	return nil
}

// DeleteRange removes bytes from start (inclusive) to end (exclusive).
func (sb *StringBuilder) DeleteRange(start, end int) error {
	if start < 0 || end > len(sb.data) || start > end {
		return ErrStringIndexOutOfRange
	}
	sb.data = append(sb.data[:start], sb.data[end:]...)
	return nil
}

// CharAt returns the byte at the given index.
func (sb *StringBuilder) CharAt(index int) (byte, error) {
	if index < 0 || index >= len(sb.data) {
		return 0, ErrStringIndexOutOfRange
	}
	return sb.data[index], nil
}

// Substring returns a new string from start (inclusive) to end (exclusive).
func (sb *StringBuilder) Substring(start, end int) (string, error) {
	if start < 0 || end > len(sb.data) || start > end {
		return "", ErrStringIndexOutOfRange
	}
	return string(sb.data[start:end]), nil
}

// IndexOf returns the index of the first occurrence of target, or -1 if not found.
func (sb *StringBuilder) IndexOf(target string) int {
	t := []byte(target)
	for i := 0; i <= len(sb.data)-len(t); i++ {
		match := true
		for j := 0; j < len(t); j++ {
			if sb.data[i+j] != t[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

// Contains returns true if the string contains the target substring.
func (sb *StringBuilder) Contains(target string) bool {
	return sb.IndexOf(target) >= 0
}

// Replace replaces the first occurrence of old with new. Returns true if replaced.
func (sb *StringBuilder) Replace(old, new string) bool {
	idx := sb.IndexOf(old)
	if idx == -1 {
		return false
	}
	result := make([]byte, 0, len(sb.data)-len(old)+len(new))
	result = append(result, sb.data[:idx]...)
	result = append(result, new...)
	result = append(result, sb.data[idx+len(old):]...)
	sb.data = result
	return true
}

// ReplaceAll replaces all occurrences of old with new. Returns the count of replacements.
func (sb *StringBuilder) ReplaceAll(old, new string) int {
	count := 0
	result := make([]byte, 0, len(sb.data))
	oldBytes := []byte(old)
	i := 0
	for i <= len(sb.data)-len(oldBytes) {
		match := true
		for j := 0; j < len(oldBytes); j++ {
			if sb.data[i+j] != oldBytes[j] {
				match = false
				break
			}
		}
		if match {
			result = append(result, new...)
			i += len(oldBytes)
			count++
		} else {
			result = append(result, sb.data[i])
			i++
		}
	}
	result = append(result, sb.data[i:]...)
	sb.data = result
	return count
}

// Reverse reverses the string in place.
func (sb *StringBuilder) Reverse() {
	for i, j := 0, len(sb.data)-1; i < j; i, j = i+1, j-1 {
		sb.data[i], sb.data[j] = sb.data[j], sb.data[i]
	}
}

// IsPalindrome returns true if the string reads the same forwards and backwards.
func (sb *StringBuilder) IsPalindrome() bool {
	for i, j := 0, len(sb.data)-1; i < j; i, j = i+1, j-1 {
		if sb.data[i] != sb.data[j] {
			return false
		}
	}
	return true
}

// Len returns the byte length of the string.
func (sb *StringBuilder) Len() int {
	return len(sb.data)
}

// String returns the string representation.
func (sb *StringBuilder) String() string {
	return string(sb.data)
}

// CountChar counts the occurrences of a byte in the string.
func (sb *StringBuilder) CountChar(ch byte) int {
	count := 0
	for _, b := range sb.data {
		if b == ch {
			count++
		}
	}
	return count
}

// ToCharArray returns a copy of the underlying bytes as a slice.
func (sb *StringBuilder) ToCharArray() []byte {
	result := make([]byte, len(sb.data))
	copy(result, sb.data)
	return result
}
