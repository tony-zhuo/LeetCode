package problems

import "testing"

func Test_MyQueue(t *testing.T) {
	t.Run("example from problem", func(t *testing.T) {
		q := Constructor()
		q.Push(1)
		q.Push(2)
		if got := q.Peek(); got != 1 {
			t.Errorf("Peek() = %v, want 1", got)
		}
		if got := q.Pop(); got != 1 {
			t.Errorf("Pop() = %v, want 1", got)
		}
		if got := q.Empty(); got != false {
			t.Errorf("Empty() = %v, want false", got)
		}
	})

	t.Run("empty check", func(t *testing.T) {
		q := Constructor()
		if got := q.Empty(); got != true {
			t.Errorf("Empty() = %v, want true", got)
		}
		q.Push(1)
		if got := q.Empty(); got != false {
			t.Errorf("Empty() = %v, want false", got)
		}
		q.Pop()
		if got := q.Empty(); got != true {
			t.Errorf("Empty() = %v, want true", got)
		}
	})

	t.Run("FIFO order", func(t *testing.T) {
		q := Constructor()
		for i := 1; i <= 5; i++ {
			q.Push(i)
		}
		for i := 1; i <= 5; i++ {
			if got := q.Pop(); got != i {
				t.Errorf("Pop() = %v, want %v", got, i)
			}
		}
	})

	t.Run("interleaved push and pop", func(t *testing.T) {
		q := Constructor()
		q.Push(1)
		q.Push(2)
		if got := q.Pop(); got != 1 {
			t.Errorf("Pop() = %v, want 1", got)
		}
		q.Push(3)
		if got := q.Pop(); got != 2 {
			t.Errorf("Pop() = %v, want 2", got)
		}
		if got := q.Pop(); got != 3 {
			t.Errorf("Pop() = %v, want 3", got)
		}
	})

	t.Run("peek does not remove", func(t *testing.T) {
		q := Constructor()
		q.Push(42)
		if got := q.Peek(); got != 42 {
			t.Errorf("Peek() = %v, want 42", got)
		}
		if got := q.Peek(); got != 42 {
			t.Errorf("Peek() = %v, want 42 (second call)", got)
		}
		if got := q.Empty(); got != false {
			t.Errorf("Empty() = %v, want false", got)
		}
	})
}
