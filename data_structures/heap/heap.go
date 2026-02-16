package datastructures

import "errors"

var ErrEmptyHeap = errors.New("heap is empty")

// MinHeap - array-based binary heap where parent <= children
type MinHeap struct {
	items []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		items: []int{},
	}
}

func (h *MinHeap) Insert(val int) {
	h.items = append(h.items, val)
	h.siftUp(len(h.items) - 1)
}

func (h *MinHeap) ExtractMin() (int, error) {
	if h.IsEmpty() {
		return 0, ErrEmptyHeap
	}

	min := h.items[0]
	last := len(h.items) - 1
	h.items[0] = h.items[last]
	h.items = h.items[:last]

	if !h.IsEmpty() {
		h.siftDown(0)
	}

	return min, nil
}

func (h *MinHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, ErrEmptyHeap
	}

	return h.items[0], nil
}

func (h *MinHeap) Size() int {
	return len(h.items)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.items) == 0
}

func (h *MinHeap) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.items[index] < h.items[parent] {
			h.items[index], h.items[parent] = h.items[parent], h.items[index]
			index = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) siftDown(index int) {
	size := len(h.items)
	for {
		smallest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < size && h.items[left] < h.items[smallest] {
			smallest = left
		}
		if right < size && h.items[right] < h.items[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		index = smallest
	}
}

// MaxHeap - array-based binary heap where parent >= children
type MaxHeap struct {
	items []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		items: []int{},
	}
}

func (h *MaxHeap) Insert(val int) {
	h.items = append(h.items, val)
	h.siftUp(len(h.items) - 1)
}

func (h *MaxHeap) ExtractMax() (int, error) {
	if h.IsEmpty() {
		return 0, ErrEmptyHeap
	}

	max := h.items[0]
	last := len(h.items) - 1
	h.items[0] = h.items[last]
	h.items = h.items[:last]

	if !h.IsEmpty() {
		h.siftDown(0)
	}

	return max, nil
}

func (h *MaxHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, ErrEmptyHeap
	}

	return h.items[0], nil
}

func (h *MaxHeap) Size() int {
	return len(h.items)
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.items) == 0
}

func (h *MaxHeap) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.items[index] > h.items[parent] {
			h.items[index], h.items[parent] = h.items[parent], h.items[index]
			index = parent
		} else {
			break
		}
	}
}

func (h *MaxHeap) siftDown(index int) {
	size := len(h.items)
	for {
		largest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < size && h.items[left] > h.items[largest] {
			largest = left
		}
		if right < size && h.items[right] > h.items[largest] {
			largest = right
		}

		if largest == index {
			break
		}

		h.items[index], h.items[largest] = h.items[largest], h.items[index]
		index = largest
	}
}

// Heapify builds a MinHeap from an unsorted slice in O(n)
func Heapify(arr []int) *MinHeap {
	items := make([]int, len(arr))
	copy(items, arr)

	h := &MinHeap{items: items}

	for i := len(items)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}

	return h
}
