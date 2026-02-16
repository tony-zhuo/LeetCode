package datastructures

import (
	"errors"
	"testing"
)

func TestStack_PushAndPopLIFO(t *testing.T) {
	tests := []struct {
		name     string
		pushVals []int
		wantPops []int
	}{
		{
			name:     "single element",
			pushVals: []int{1},
			wantPops: []int{1},
		},
		{
			name:     "multiple elements in LIFO order",
			pushVals: []int{1, 2, 3, 4, 5},
			wantPops: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "negative and zero values",
			pushVals: []int{-3, 0, 7},
			wantPops: []int{7, 0, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			for _, v := range tt.pushVals {
				s.Push(v)
			}

			for i, want := range tt.wantPops {
				got, err := s.Pop()
				if err != nil {
					t.Fatalf("Pop() #%d unexpected error: %v", i, err)
				}
				if got != want {
					t.Errorf("Pop() #%d = %d, want %d", i, got, want)
				}
			}
		})
	}
}

func TestStack_PopEmptyStack(t *testing.T) {
	tests := []struct {
		name string
		setup func() *Stack
	}{
		{
			name: "pop on brand new empty stack",
			setup: func() *Stack {
				return NewStack()
			},
		},
		{
			name: "pop after all elements removed",
			setup: func() *Stack {
				s := NewStack()
				s.Push(10)
				s.Push(20)
				s.Pop()
				s.Pop()
				return s
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.setup()
			_, err := s.Pop()
			if !errors.Is(err, ErrEmptyStack) {
				t.Errorf("Pop() error = %v, want %v", err, ErrEmptyStack)
			}
		})
	}
}

func TestStack_PeekEmptyStack(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *Stack
	}{
		{
			name: "peek on brand new empty stack",
			setup: func() *Stack {
				return NewStack()
			},
		},
		{
			name: "peek after all elements removed",
			setup: func() *Stack {
				s := NewStack()
				s.Push(42)
				s.Pop()
				return s
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.setup()
			_, err := s.Peek()
			if !errors.Is(err, ErrEmptyStack) {
				t.Errorf("Peek() error = %v, want %v", err, ErrEmptyStack)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name     string
		pushVals []int
		want     int
	}{
		{
			name:     "peek returns top without removing",
			pushVals: []int{1, 2, 3},
			want:     3,
		},
		{
			name:     "peek single element",
			pushVals: []int{99},
			want:     99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			for _, v := range tt.pushVals {
				s.Push(v)
			}

			got, err := s.Peek()
			if err != nil {
				t.Fatalf("Peek() unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Peek() = %d, want %d", got, tt.want)
			}

			// Verify Peek does not remove the element
			if s.Size() != len(tt.pushVals) {
				t.Errorf("Size() after Peek() = %d, want %d", s.Size(), len(tt.pushVals))
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name       string
		operations func(s *Stack)
		wantSize   int
	}{
		{
			name:       "empty stack has size 0",
			operations: func(s *Stack) {},
			wantSize:   0,
		},
		{
			name: "size after pushes",
			operations: func(s *Stack) {
				s.Push(1)
				s.Push(2)
				s.Push(3)
			},
			wantSize: 3,
		},
		{
			name: "size after mixed push and pop",
			operations: func(s *Stack) {
				s.Push(1)
				s.Push(2)
				s.Push(3)
				s.Pop()
				s.Push(4)
				s.Pop()
				s.Pop()
			},
			wantSize: 1,
		},
		{
			name: "size returns to 0 after removing all",
			operations: func(s *Stack) {
				s.Push(10)
				s.Push(20)
				s.Pop()
				s.Pop()
			},
			wantSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			tt.operations(s)
			if got := s.Size(); got != tt.wantSize {
				t.Errorf("Size() = %d, want %d", got, tt.wantSize)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name       string
		operations func(s *Stack)
		want       bool
	}{
		{
			name:       "new stack is empty",
			operations: func(s *Stack) {},
			want:       true,
		},
		{
			name: "stack with elements is not empty",
			operations: func(s *Stack) {
				s.Push(1)
			},
			want: false,
		},
		{
			name: "stack is empty after removing all elements",
			operations: func(s *Stack) {
				s.Push(1)
				s.Push(2)
				s.Pop()
				s.Pop()
			},
			want: true,
		},
		{
			name: "stack is not empty after push-pop-push",
			operations: func(s *Stack) {
				s.Push(5)
				s.Pop()
				s.Push(10)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			tt.operations(s)
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
