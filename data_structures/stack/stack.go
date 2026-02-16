package datastructures

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return top, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}
