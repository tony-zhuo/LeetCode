package datastructures

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")
var ErrFullQueue = errors.New("queue is full")

// Queue is a standard slice-backed FIFO queue.
type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, ErrEmptyQueue
	}

	front := q.items[0]
	q.items = q.items[1:]

	return front, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, ErrEmptyQueue
	}

	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

// CircularQueue is a fixed-size ring buffer FIFO queue.
type CircularQueue struct {
	items []int
	head  int
	tail  int
	size  int
	cap   int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, capacity),
		head:  0,
		tail:  0,
		size:  0,
		cap:   capacity,
	}
}

func (q *CircularQueue) Enqueue(val int) error {
	if q.IsFull() {
		return ErrFullQueue
	}

	q.items[q.tail] = val
	q.tail = (q.tail + 1) % q.cap
	q.size++

	return nil
}

func (q *CircularQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, ErrEmptyQueue
	}

	front := q.items[q.head]
	q.head = (q.head + 1) % q.cap
	q.size--

	return front, nil
}

func (q *CircularQueue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, ErrEmptyQueue
	}

	return q.items[q.head], nil
}

func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue) IsFull() bool {
	return q.size == q.cap
}

func (q *CircularQueue) Size() int {
	return q.size
}
