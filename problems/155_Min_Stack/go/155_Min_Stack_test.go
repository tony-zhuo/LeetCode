package problems

import "testing"

func Test_MinStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []int // -1 means no return value (push/pop)
	}{
		{
			name:       "example 1",
			operations: []string{"push", "push", "push", "getMin", "pop", "top", "getMin"},
			values:     []int{-2, 0, -3, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, -3, -1, 0, -2},
		},
		{
			name:       "single element",
			operations: []string{"push", "top", "getMin", "pop"},
			values:     []int{42, 0, 0, 0},
			expected:   []int{-1, 42, 42, -1},
		},
		{
			name:       "push duplicates then pop",
			operations: []string{"push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			values:     []int{1, 1, 1, 0, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, 1, -1, 1, -1, 1},
		},
		{
			name:       "min changes after pop",
			operations: []string{"push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			values:     []int{3, 1, 2, 0, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, 1, -1, 1, -1, 3},
		},
		{
			name:       "decreasing order",
			operations: []string{"push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			values:     []int{3, 2, 1, 0, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, 1, -1, 2, -1, 3},
		},
		{
			name:       "negative values",
			operations: []string{"push", "push", "getMin", "pop", "getMin"},
			values:     []int{-1, -2, 0, 0, 0},
			expected:   []int{-1, -1, -2, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor()
			for i, op := range tt.operations {
				switch op {
				case "push":
					obj.Push(tt.values[i])
				case "pop":
					obj.Pop()
				case "top":
					got := obj.Top()
					if got != tt.expected[i] {
						t.Errorf("step %d top() = %d, want %d", i, got, tt.expected[i])
					}
				case "getMin":
					got := obj.GetMin()
					if got != tt.expected[i] {
						t.Errorf("step %d getMin() = %d, want %d", i, got, tt.expected[i])
					}
				}
			}
		})
	}
}