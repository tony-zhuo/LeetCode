package datastructures

import (
	"errors"
	"sort"
	"testing"
)

func TestMinHeap_InsertAndExtractMin(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		wantOrder  []int
	}{
		{
			name:       "single element",
			insertVals: []int{42},
			wantOrder:  []int{42},
		},
		{
			name:       "ascending input",
			insertVals: []int{1, 2, 3, 4, 5},
			wantOrder:  []int{1, 2, 3, 4, 5},
		},
		{
			name:       "descending input",
			insertVals: []int{5, 4, 3, 2, 1},
			wantOrder:  []int{1, 2, 3, 4, 5},
		},
		{
			name:       "random sequence",
			insertVals: []int{38, 7, 15, 2, 44, 9, 1, 23},
			wantOrder:  []int{1, 2, 7, 9, 15, 23, 38, 44},
		},
		{
			name:       "duplicates",
			insertVals: []int{5, 3, 5, 1, 3, 1},
			wantOrder:  []int{1, 1, 3, 3, 5, 5},
		},
		{
			name:       "negative values",
			insertVals: []int{-10, 0, 5, -3, 8, -1},
			wantOrder:  []int{-10, -3, -1, 0, 5, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, v := range tt.insertVals {
				h.Insert(v)
			}

			if h.Size() != len(tt.insertVals) {
				t.Fatalf("Size() = %d, want %d", h.Size(), len(tt.insertVals))
			}

			for i, want := range tt.wantOrder {
				got, err := h.ExtractMin()
				if err != nil {
					t.Fatalf("ExtractMin() #%d unexpected error: %v", i, err)
				}
				if got != want {
					t.Errorf("ExtractMin() #%d = %d, want %d", i, got, want)
				}
			}

			if !h.IsEmpty() {
				t.Errorf("heap should be empty after extracting all elements, Size() = %d", h.Size())
			}
		})
	}
}

func TestMaxHeap_InsertAndExtractMax(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		wantOrder  []int
	}{
		{
			name:       "single element",
			insertVals: []int{42},
			wantOrder:  []int{42},
		},
		{
			name:       "ascending input",
			insertVals: []int{1, 2, 3, 4, 5},
			wantOrder:  []int{5, 4, 3, 2, 1},
		},
		{
			name:       "descending input",
			insertVals: []int{5, 4, 3, 2, 1},
			wantOrder:  []int{5, 4, 3, 2, 1},
		},
		{
			name:       "random sequence",
			insertVals: []int{38, 7, 15, 2, 44, 9, 1, 23},
			wantOrder:  []int{44, 38, 23, 15, 9, 7, 2, 1},
		},
		{
			name:       "duplicates",
			insertVals: []int{5, 3, 5, 1, 3, 1},
			wantOrder:  []int{5, 5, 3, 3, 1, 1},
		},
		{
			name:       "negative values",
			insertVals: []int{-10, 0, 5, -3, 8, -1},
			wantOrder:  []int{8, 5, 0, -1, -3, -10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMaxHeap()
			for _, v := range tt.insertVals {
				h.Insert(v)
			}

			if h.Size() != len(tt.insertVals) {
				t.Fatalf("Size() = %d, want %d", h.Size(), len(tt.insertVals))
			}

			for i, want := range tt.wantOrder {
				got, err := h.ExtractMax()
				if err != nil {
					t.Fatalf("ExtractMax() #%d unexpected error: %v", i, err)
				}
				if got != want {
					t.Errorf("ExtractMax() #%d = %d, want %d", i, got, want)
				}
			}

			if !h.IsEmpty() {
				t.Errorf("heap should be empty after extracting all elements, Size() = %d", h.Size())
			}
		})
	}
}

func TestMinHeap_PeekDoesNotModifyHeap(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		wantPeek   int
	}{
		{
			name:       "peek single element",
			insertVals: []int{10},
			wantPeek:   10,
		},
		{
			name:       "peek multiple elements",
			insertVals: []int{30, 10, 20},
			wantPeek:   10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, v := range tt.insertVals {
				h.Insert(v)
			}

			sizeBefore := h.Size()

			got, err := h.Peek()
			if err != nil {
				t.Fatalf("Peek() unexpected error: %v", err)
			}
			if got != tt.wantPeek {
				t.Errorf("Peek() = %d, want %d", got, tt.wantPeek)
			}
			if h.Size() != sizeBefore {
				t.Errorf("Size() after Peek() = %d, want %d", h.Size(), sizeBefore)
			}

			// Peek again to confirm stability
			got2, err := h.Peek()
			if err != nil {
				t.Fatalf("Peek() second call unexpected error: %v", err)
			}
			if got2 != tt.wantPeek {
				t.Errorf("Peek() second call = %d, want %d", got2, tt.wantPeek)
			}
		})
	}
}

func TestMaxHeap_PeekDoesNotModifyHeap(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		wantPeek   int
	}{
		{
			name:       "peek single element",
			insertVals: []int{10},
			wantPeek:   10,
		},
		{
			name:       "peek multiple elements",
			insertVals: []int{10, 30, 20},
			wantPeek:   30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMaxHeap()
			for _, v := range tt.insertVals {
				h.Insert(v)
			}

			sizeBefore := h.Size()

			got, err := h.Peek()
			if err != nil {
				t.Fatalf("Peek() unexpected error: %v", err)
			}
			if got != tt.wantPeek {
				t.Errorf("Peek() = %d, want %d", got, tt.wantPeek)
			}
			if h.Size() != sizeBefore {
				t.Errorf("Size() after Peek() = %d, want %d", h.Size(), sizeBefore)
			}
		})
	}
}

func TestMinHeap_ExtractFromEmptyHeap(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *MinHeap
	}{
		{
			name: "extract from brand new empty heap",
			setup: func() *MinHeap {
				return NewMinHeap()
			},
		},
		{
			name: "extract after all elements removed",
			setup: func() *MinHeap {
				h := NewMinHeap()
				h.Insert(10)
				h.Insert(20)
				h.ExtractMin()
				h.ExtractMin()
				return h
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.setup()
			_, err := h.ExtractMin()
			if !errors.Is(err, ErrEmptyHeap) {
				t.Errorf("ExtractMin() error = %v, want %v", err, ErrEmptyHeap)
			}
		})
	}
}

func TestMaxHeap_ExtractFromEmptyHeap(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *MaxHeap
	}{
		{
			name: "extract from brand new empty heap",
			setup: func() *MaxHeap {
				return NewMaxHeap()
			},
		},
		{
			name: "extract after all elements removed",
			setup: func() *MaxHeap {
				h := NewMaxHeap()
				h.Insert(10)
				h.Insert(20)
				h.ExtractMax()
				h.ExtractMax()
				return h
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.setup()
			_, err := h.ExtractMax()
			if !errors.Is(err, ErrEmptyHeap) {
				t.Errorf("ExtractMax() error = %v, want %v", err, ErrEmptyHeap)
			}
		})
	}
}

func TestMinHeap_PeekEmptyHeap(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *MinHeap
	}{
		{
			name: "peek on brand new empty heap",
			setup: func() *MinHeap {
				return NewMinHeap()
			},
		},
		{
			name: "peek after all elements removed",
			setup: func() *MinHeap {
				h := NewMinHeap()
				h.Insert(5)
				h.ExtractMin()
				return h
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.setup()
			_, err := h.Peek()
			if !errors.Is(err, ErrEmptyHeap) {
				t.Errorf("Peek() error = %v, want %v", err, ErrEmptyHeap)
			}
		})
	}
}

func TestMaxHeap_PeekEmptyHeap(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *MaxHeap
	}{
		{
			name: "peek on brand new empty heap",
			setup: func() *MaxHeap {
				return NewMaxHeap()
			},
		},
		{
			name: "peek after all elements removed",
			setup: func() *MaxHeap {
				h := NewMaxHeap()
				h.Insert(5)
				h.ExtractMax()
				return h
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.setup()
			_, err := h.Peek()
			if !errors.Is(err, ErrEmptyHeap) {
				t.Errorf("Peek() error = %v, want %v", err, ErrEmptyHeap)
			}
		})
	}
}

func TestHeapify(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantOrder []int
	}{
		{
			name:      "empty slice",
			input:     []int{},
			wantOrder: []int{},
		},
		{
			name:      "single element",
			input:     []int{42},
			wantOrder: []int{42},
		},
		{
			name:      "already sorted",
			input:     []int{1, 2, 3, 4, 5},
			wantOrder: []int{1, 2, 3, 4, 5},
		},
		{
			name:      "reverse sorted",
			input:     []int{5, 4, 3, 2, 1},
			wantOrder: []int{1, 2, 3, 4, 5},
		},
		{
			name:      "random order",
			input:     []int{15, 3, 9, 1, 22, 7, 11, 4},
			wantOrder: []int{1, 3, 4, 7, 9, 11, 15, 22},
		},
		{
			name:      "duplicates",
			input:     []int{5, 1, 5, 1, 3},
			wantOrder: []int{1, 1, 3, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Heapify(tt.input)

			if h.Size() != len(tt.input) {
				t.Fatalf("Size() = %d, want %d", h.Size(), len(tt.input))
			}

			for i, want := range tt.wantOrder {
				got, err := h.ExtractMin()
				if err != nil {
					t.Fatalf("ExtractMin() #%d unexpected error: %v", i, err)
				}
				if got != want {
					t.Errorf("ExtractMin() #%d = %d, want %d", i, got, want)
				}
			}
		})
	}
}

func TestHeapify_DoesNotModifyOriginalSlice(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	backup := make([]int, len(original))
	copy(backup, original)

	Heapify(original)

	for i := range original {
		if original[i] != backup[i] {
			t.Errorf("original slice modified at index %d: got %d, want %d", i, original[i], backup[i])
		}
	}
}

func TestMinHeap_Size(t *testing.T) {
	tests := []struct {
		name       string
		operations func(h *MinHeap)
		wantSize   int
	}{
		{
			name:       "empty heap has size 0",
			operations: func(h *MinHeap) {},
			wantSize:   0,
		},
		{
			name: "size after inserts",
			operations: func(h *MinHeap) {
				h.Insert(1)
				h.Insert(2)
				h.Insert(3)
			},
			wantSize: 3,
		},
		{
			name: "size after mixed insert and extract",
			operations: func(h *MinHeap) {
				h.Insert(1)
				h.Insert(2)
				h.Insert(3)
				h.ExtractMin()
				h.Insert(4)
				h.ExtractMin()
			},
			wantSize: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			tt.operations(h)
			if got := h.Size(); got != tt.wantSize {
				t.Errorf("Size() = %d, want %d", got, tt.wantSize)
			}
		})
	}
}

func TestMinHeap_IsEmpty(t *testing.T) {
	tests := []struct {
		name       string
		operations func(h *MinHeap)
		want       bool
	}{
		{
			name:       "new heap is empty",
			operations: func(h *MinHeap) {},
			want:       true,
		},
		{
			name: "heap with elements is not empty",
			operations: func(h *MinHeap) {
				h.Insert(1)
			},
			want: false,
		},
		{
			name: "heap is empty after removing all elements",
			operations: func(h *MinHeap) {
				h.Insert(1)
				h.Insert(2)
				h.ExtractMin()
				h.ExtractMin()
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			tt.operations(h)
			if got := h.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_LargeDataset(t *testing.T) {
	h := NewMinHeap()
	vals := []int{100, 50, 200, 25, 75, 150, 250, 10, 30, 60, 80, 125, 175, 225, 300}

	for _, v := range vals {
		h.Insert(v)
	}

	sorted := make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)

	for i, want := range sorted {
		got, err := h.ExtractMin()
		if err != nil {
			t.Fatalf("ExtractMin() #%d unexpected error: %v", i, err)
		}
		if got != want {
			t.Errorf("ExtractMin() #%d = %d, want %d", i, got, want)
		}
	}
}
