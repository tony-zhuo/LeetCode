package datastructures

import (
	"math"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	t.Run("query full range equals sum of all elements", func(t *testing.T) {
		nums := []int{1, 3, 5, 7, 9, 11}
		st := NewSegmentTree(nums)
		got := st.Query(0, 5)
		want := 1 + 3 + 5 + 7 + 9 + 11
		if got != want {
			t.Errorf("Query(0, 5) = %d, want %d", got, want)
		}
	})

	t.Run("query sub-ranges", func(t *testing.T) {
		nums := []int{1, 3, 5, 7, 9, 11}
		st := NewSegmentTree(nums)

		tests := []struct {
			name       string
			left       int
			right      int
			want       int
		}{
			{"first half", 0, 2, 1 + 3 + 5},
			{"second half", 3, 5, 7 + 9 + 11},
			{"middle range", 1, 4, 3 + 5 + 7 + 9},
			{"two elements", 2, 3, 5 + 7},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := st.Query(tt.left, tt.right)
				if got != tt.want {
					t.Errorf("Query(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
				}
			})
		}
	})

	t.Run("update element and re-query", func(t *testing.T) {
		nums := []int{1, 3, 5, 7, 9, 11}
		st := NewSegmentTree(nums)

		// Update index 2 from 5 to 10
		st.Update(2, 10)

		got := st.Query(0, 5)
		want := 1 + 3 + 10 + 7 + 9 + 11
		if got != want {
			t.Errorf("after Update(2, 10), Query(0, 5) = %d, want %d", got, want)
		}

		// Verify sub-range containing updated element
		got = st.Query(1, 3)
		want = 3 + 10 + 7
		if got != want {
			t.Errorf("after Update(2, 10), Query(1, 3) = %d, want %d", got, want)
		}

		// Verify sub-range not containing updated element
		got = st.Query(3, 5)
		want = 7 + 9 + 11
		if got != want {
			t.Errorf("after Update(2, 10), Query(3, 5) = %d, want %d", got, want)
		}
	})

	t.Run("single element query", func(t *testing.T) {
		nums := []int{1, 3, 5, 7, 9, 11}
		st := NewSegmentTree(nums)

		for i, v := range nums {
			got := st.Query(i, i)
			if got != v {
				t.Errorf("Query(%d, %d) = %d, want %d", i, i, got, v)
			}
		}
	})

	t.Run("single element array", func(t *testing.T) {
		nums := []int{42}
		st := NewSegmentTree(nums)

		got := st.Query(0, 0)
		if got != 42 {
			t.Errorf("Query(0, 0) = %d, want 42", got)
		}

		st.Update(0, 100)
		got = st.Query(0, 0)
		if got != 100 {
			t.Errorf("after Update(0, 100), Query(0, 0) = %d, want 100", got)
		}
	})

	t.Run("multiple updates", func(t *testing.T) {
		nums := []int{2, 4, 6, 8}
		st := NewSegmentTree(nums)

		st.Update(0, 1)
		st.Update(3, 10)

		got := st.Query(0, 3)
		want := 1 + 4 + 6 + 10
		if got != want {
			t.Errorf("after updates, Query(0, 3) = %d, want %d", got, want)
		}
	})
}

func TestRangeMinSegmentTree(t *testing.T) {
	t.Run("query full range equals minimum element", func(t *testing.T) {
		nums := []int{5, 2, 8, 1, 9, 3}
		st := NewRangeMinSegmentTree(nums)
		got := st.Query(0, 5)
		want := 1
		if got != want {
			t.Errorf("Query(0, 5) = %d, want %d", got, want)
		}
	})

	t.Run("query sub-ranges", func(t *testing.T) {
		nums := []int{5, 2, 8, 1, 9, 3}
		st := NewRangeMinSegmentTree(nums)

		tests := []struct {
			name  string
			left  int
			right int
			want  int
		}{
			{"first half", 0, 2, 2},
			{"second half", 3, 5, 1},
			{"middle range", 1, 4, 1},
			{"last two", 4, 5, 3},
			{"first two", 0, 1, 2},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := st.Query(tt.left, tt.right)
				if got != tt.want {
					t.Errorf("Query(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
				}
			})
		}
	})

	t.Run("update element to smaller value", func(t *testing.T) {
		nums := []int{5, 2, 8, 1, 9, 3}
		st := NewRangeMinSegmentTree(nums)

		// Update index 4 from 9 to 0 (new global minimum)
		st.Update(4, 0)

		got := st.Query(0, 5)
		want := 0
		if got != want {
			t.Errorf("after Update(4, 0), Query(0, 5) = %d, want %d", got, want)
		}

		// Sub-range containing updated element
		got = st.Query(3, 5)
		want = 0
		if got != want {
			t.Errorf("after Update(4, 0), Query(3, 5) = %d, want %d", got, want)
		}
	})

	t.Run("update element to larger value changes min", func(t *testing.T) {
		nums := []int{5, 2, 8, 1, 9, 3}
		st := NewRangeMinSegmentTree(nums)

		// Update the current minimum (index 3, value 1) to a larger value
		st.Update(3, 100)

		// New global minimum should be 2 (at index 1)
		got := st.Query(0, 5)
		want := 2
		if got != want {
			t.Errorf("after Update(3, 100), Query(0, 5) = %d, want %d", got, want)
		}

		// The sub-range [3,5] should now have min = 3
		got = st.Query(3, 5)
		want = 3
		if got != want {
			t.Errorf("after Update(3, 100), Query(3, 5) = %d, want %d", got, want)
		}
	})

	t.Run("single element query", func(t *testing.T) {
		nums := []int{5, 2, 8, 1, 9, 3}
		st := NewRangeMinSegmentTree(nums)

		for i, v := range nums {
			got := st.Query(i, i)
			if got != v {
				t.Errorf("Query(%d, %d) = %d, want %d", i, i, got, v)
			}
		}
	})

	t.Run("single element array", func(t *testing.T) {
		nums := []int{7}
		st := NewRangeMinSegmentTree(nums)

		got := st.Query(0, 0)
		if got != 7 {
			t.Errorf("Query(0, 0) = %d, want 7", got)
		}

		st.Update(0, 3)
		got = st.Query(0, 0)
		if got != 3 {
			t.Errorf("after Update(0, 3), Query(0, 0) = %d, want 3", got)
		}
	})

	t.Run("out of range query returns MaxInt", func(t *testing.T) {
		nums := []int{5, 2, 8}
		st := NewRangeMinSegmentTree(nums)

		// Query a range that doesn't overlap should return math.MaxInt
		// This tests the edge case of the recursive query
		got := st.query(1, 0, 2, 3, 5)
		if got != math.MaxInt {
			t.Errorf("query with no overlap = %d, want math.MaxInt", got)
		}
	})
}
