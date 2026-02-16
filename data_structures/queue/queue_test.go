package datastructures

import "testing"

func Test_Queue_Enqueue_Dequeue_FIFO(t *testing.T) {
	tests := []struct {
		name     string
		enqueue  []int
		wantFIFO []int
	}{
		{
			name:     "single element",
			enqueue:  []int{1},
			wantFIFO: []int{1},
		},
		{
			name:     "multiple elements preserve FIFO order",
			enqueue:  []int{1, 2, 3, 4, 5},
			wantFIFO: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, v := range tt.enqueue {
				q.Enqueue(v)
			}

			for i, want := range tt.wantFIFO {
				got, err := q.Dequeue()
				if err != nil {
					t.Fatalf("Dequeue() at index %d returned unexpected error: %v", i, err)
				}
				if got != want {
					t.Errorf("Dequeue() at index %d = %d, want %d", i, got, want)
				}
			}
		})
	}
}

func Test_Queue_Dequeue_Empty(t *testing.T) {
	q := NewQueue()
	_, err := q.Dequeue()
	if err != ErrEmptyQueue {
		t.Errorf("Dequeue() on empty queue error = %v, want %v", err, ErrEmptyQueue)
	}
}

func Test_Queue_Peek_Empty(t *testing.T) {
	q := NewQueue()
	_, err := q.Peek()
	if err != ErrEmptyQueue {
		t.Errorf("Peek() on empty queue error = %v, want %v", err, ErrEmptyQueue)
	}
}

func Test_Queue_Peek(t *testing.T) {
	tests := []struct {
		name    string
		enqueue []int
		want    int
	}{
		{
			name:    "peek returns front element",
			enqueue: []int{10, 20, 30},
			want:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, v := range tt.enqueue {
				q.Enqueue(v)
			}

			got, err := q.Peek()
			if err != nil {
				t.Fatalf("Peek() returned unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Peek() = %d, want %d", got, tt.want)
			}

			// Peek should not remove the element
			if q.Size() != len(tt.enqueue) {
				t.Errorf("Size() after Peek() = %d, want %d", q.Size(), len(tt.enqueue))
			}
		})
	}
}

func Test_Queue_IsEmpty(t *testing.T) {
	tests := []struct {
		name    string
		enqueue []int
		want    bool
	}{
		{
			name:    "new queue is empty",
			enqueue: []int{},
			want:    true,
		},
		{
			name:    "queue with elements is not empty",
			enqueue: []int{1},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, v := range tt.enqueue {
				q.Enqueue(v)
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Queue_Size(t *testing.T) {
	tests := []struct {
		name    string
		enqueue []int
		want    int
	}{
		{
			name:    "empty queue has size 0",
			enqueue: []int{},
			want:    0,
		},
		{
			name:    "queue with 3 elements has size 3",
			enqueue: []int{1, 2, 3},
			want:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, v := range tt.enqueue {
				q.Enqueue(v)
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Size() = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_CircularQueue_Enqueue_Dequeue_FIFO(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		enqueue  []int
		wantFIFO []int
	}{
		{
			name:     "single element",
			capacity: 5,
			enqueue:  []int{1},
			wantFIFO: []int{1},
		},
		{
			name:     "fill to capacity",
			capacity: 3,
			enqueue:  []int{1, 2, 3},
			wantFIFO: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewCircularQueue(tt.capacity)
			for _, v := range tt.enqueue {
				if err := q.Enqueue(v); err != nil {
					t.Fatalf("Enqueue(%d) returned unexpected error: %v", v, err)
				}
			}

			for i, want := range tt.wantFIFO {
				got, err := q.Dequeue()
				if err != nil {
					t.Fatalf("Dequeue() at index %d returned unexpected error: %v", i, err)
				}
				if got != want {
					t.Errorf("Dequeue() at index %d = %d, want %d", i, got, want)
				}
			}
		})
	}
}

func Test_CircularQueue_WrapAround(t *testing.T) {
	t.Run("enqueue to full, dequeue some, enqueue more", func(t *testing.T) {
		q := NewCircularQueue(3)

		// Fill the queue: [1, 2, 3]
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		if !q.IsFull() {
			t.Fatal("expected queue to be full")
		}

		// Dequeue two elements: head wraps forward
		got, _ := q.Dequeue()
		if got != 1 {
			t.Errorf("Dequeue() = %d, want 1", got)
		}
		got, _ = q.Dequeue()
		if got != 2 {
			t.Errorf("Dequeue() = %d, want 2", got)
		}

		// Enqueue two more, causing tail to wrap around: [4, 5, 3] internally
		q.Enqueue(4)
		q.Enqueue(5)

		if !q.IsFull() {
			t.Fatal("expected queue to be full after wrap-around enqueue")
		}

		// Dequeue all remaining in FIFO order: 3, 4, 5
		expected := []int{3, 4, 5}
		for i, want := range expected {
			got, err := q.Dequeue()
			if err != nil {
				t.Fatalf("Dequeue() at index %d returned unexpected error: %v", i, err)
			}
			if got != want {
				t.Errorf("Dequeue() at index %d = %d, want %d", i, got, want)
			}
		}

		if !q.IsEmpty() {
			t.Error("expected queue to be empty after dequeuing all elements")
		}
	})
}

func Test_CircularQueue_Full_Rejection(t *testing.T) {
	q := NewCircularQueue(2)
	q.Enqueue(1)
	q.Enqueue(2)

	err := q.Enqueue(3)
	if err != ErrFullQueue {
		t.Errorf("Enqueue() on full queue error = %v, want %v", err, ErrFullQueue)
	}
}

func Test_CircularQueue_Dequeue_Empty(t *testing.T) {
	q := NewCircularQueue(3)
	_, err := q.Dequeue()
	if err != ErrEmptyQueue {
		t.Errorf("Dequeue() on empty queue error = %v, want %v", err, ErrEmptyQueue)
	}
}

func Test_CircularQueue_Peek_Empty(t *testing.T) {
	q := NewCircularQueue(3)
	_, err := q.Peek()
	if err != ErrEmptyQueue {
		t.Errorf("Peek() on empty queue error = %v, want %v", err, ErrEmptyQueue)
	}
}

func Test_CircularQueue_Peek(t *testing.T) {
	q := NewCircularQueue(3)
	q.Enqueue(10)
	q.Enqueue(20)

	got, err := q.Peek()
	if err != nil {
		t.Fatalf("Peek() returned unexpected error: %v", err)
	}
	if got != 10 {
		t.Errorf("Peek() = %d, want 10", got)
	}

	// Peek should not remove the element
	if q.Size() != 2 {
		t.Errorf("Size() after Peek() = %d, want 2", q.Size())
	}
}

func Test_CircularQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		enqueue  []int
		want     bool
	}{
		{
			name:     "new circular queue is empty",
			capacity: 3,
			enqueue:  []int{},
			want:     true,
		},
		{
			name:     "circular queue with elements is not empty",
			capacity: 3,
			enqueue:  []int{1},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewCircularQueue(tt.capacity)
			for _, v := range tt.enqueue {
				q.Enqueue(v)
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CircularQueue_IsFull(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		enqueue  []int
		want     bool
	}{
		{
			name:     "not full when partially filled",
			capacity: 3,
			enqueue:  []int{1},
			want:     false,
		},
		{
			name:     "full when filled to capacity",
			capacity: 2,
			enqueue:  []int{1, 2},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewCircularQueue(tt.capacity)
			for _, v := range tt.enqueue {
				q.Enqueue(v)
			}
			if got := q.IsFull(); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CircularQueue_Size(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		enqueue  []int
		want     int
	}{
		{
			name:     "empty circular queue has size 0",
			capacity: 3,
			enqueue:  []int{},
			want:     0,
		},
		{
			name:     "circular queue with 2 elements has size 2",
			capacity: 5,
			enqueue:  []int{1, 2},
			want:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewCircularQueue(tt.capacity)
			for _, v := range tt.enqueue {
				q.Enqueue(v)
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Size() = %d, want %d", got, tt.want)
			}
		})
	}
}
