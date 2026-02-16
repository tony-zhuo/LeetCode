package datastructures

import (
	"reflect"
	"testing"
)

// ============================================================
// Singly Linked List Tests
// ============================================================

func TestSinglyLinkedList_InsertAtHead(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "insert single element",
			values: []int{1},
			want:   []int{1},
		},
		{
			name:   "insert multiple elements",
			values: []int{1, 2, 3},
			want:   []int{3, 2, 1},
		},
		{
			name:   "insert no elements",
			values: []int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.values {
				l.InsertAtHead(v)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
			if l.Len() != len(tt.want) {
				t.Errorf("Len() = %d, want %d", l.Len(), len(tt.want))
			}
		})
	}
}

func TestSinglyLinkedList_InsertAtTail(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "insert single element",
			values: []int{1},
			want:   []int{1},
		},
		{
			name:   "insert multiple elements",
			values: []int{1, 2, 3},
			want:   []int{1, 2, 3},
		},
		{
			name:   "insert no elements",
			values: []int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.values {
				l.InsertAtTail(v)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
			if l.Len() != len(tt.want) {
				t.Errorf("Len() = %d, want %d", l.Len(), len(tt.want))
			}
		})
	}
}

func TestSinglyLinkedList_InsertAt(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		index     int
		val       int
		want      []int
		wantErr   error
	}{
		{
			name:    "insert at beginning of empty list",
			initial: []int{},
			index:   0,
			val:     1,
			want:    []int{1},
			wantErr: nil,
		},
		{
			name:    "insert at beginning",
			initial: []int{2, 3},
			index:   0,
			val:     1,
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "insert in middle",
			initial: []int{1, 3},
			index:   1,
			val:     2,
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "insert at end",
			initial: []int{1, 2},
			index:   2,
			val:     3,
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "insert at negative index",
			initial: []int{1, 2},
			index:   -1,
			val:     0,
			want:    []int{1, 2},
			wantErr: ErrIndexOutOfRange,
		},
		{
			name:    "insert beyond size",
			initial: []int{1, 2},
			index:   3,
			val:     4,
			want:    []int{1, 2},
			wantErr: ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			err := l.InsertAt(tt.index, tt.val)
			if err != tt.wantErr {
				t.Errorf("InsertAt() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_DeleteAtHead(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		wantVal int
		wantErr error
		wantList []int
	}{
		{
			name:     "delete from empty list",
			initial:  []int{},
			wantVal:  0,
			wantErr:  ErrEmptyList,
			wantList: []int{},
		},
		{
			name:     "delete single element",
			initial:  []int{1},
			wantVal:  1,
			wantErr:  nil,
			wantList: []int{},
		},
		{
			name:     "delete from multiple elements",
			initial:  []int{1, 2, 3},
			wantVal:  1,
			wantErr:  nil,
			wantList: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			val, err := l.DeleteAtHead()
			if err != tt.wantErr {
				t.Errorf("DeleteAtHead() error = %v, wantErr %v", err, tt.wantErr)
			}
			if val != tt.wantVal {
				t.Errorf("DeleteAtHead() val = %d, want %d", val, tt.wantVal)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.wantList) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.wantList)
			}
		})
	}
}

func TestSinglyLinkedList_DeleteAtTail(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		wantVal  int
		wantErr  error
		wantList []int
	}{
		{
			name:     "delete from empty list",
			initial:  []int{},
			wantVal:  0,
			wantErr:  ErrEmptyList,
			wantList: []int{},
		},
		{
			name:     "delete single element",
			initial:  []int{1},
			wantVal:  1,
			wantErr:  nil,
			wantList: []int{},
		},
		{
			name:     "delete from multiple elements",
			initial:  []int{1, 2, 3},
			wantVal:  3,
			wantErr:  nil,
			wantList: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			val, err := l.DeleteAtTail()
			if err != tt.wantErr {
				t.Errorf("DeleteAtTail() error = %v, wantErr %v", err, tt.wantErr)
			}
			if val != tt.wantVal {
				t.Errorf("DeleteAtTail() val = %d, want %d", val, tt.wantVal)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.wantList) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.wantList)
			}
		})
	}
}

func TestSinglyLinkedList_DeleteAt(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		index    int
		wantVal  int
		wantErr  error
		wantList []int
	}{
		{
			name:     "delete from empty list",
			initial:  []int{},
			index:    0,
			wantVal:  0,
			wantErr:  ErrEmptyList,
			wantList: []int{},
		},
		{
			name:     "delete at beginning",
			initial:  []int{1, 2, 3},
			index:    0,
			wantVal:  1,
			wantErr:  nil,
			wantList: []int{2, 3},
		},
		{
			name:     "delete in middle",
			initial:  []int{1, 2, 3},
			index:    1,
			wantVal:  2,
			wantErr:  nil,
			wantList: []int{1, 3},
		},
		{
			name:     "delete at end",
			initial:  []int{1, 2, 3},
			index:    2,
			wantVal:  3,
			wantErr:  nil,
			wantList: []int{1, 2},
		},
		{
			name:     "delete at negative index",
			initial:  []int{1, 2},
			index:    -1,
			wantVal:  0,
			wantErr:  ErrIndexOutOfRange,
			wantList: []int{1, 2},
		},
		{
			name:     "delete beyond size",
			initial:  []int{1, 2},
			index:    2,
			wantVal:  0,
			wantErr:  ErrIndexOutOfRange,
			wantList: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			val, err := l.DeleteAt(tt.index)
			if err != tt.wantErr {
				t.Errorf("DeleteAt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if val != tt.wantVal {
				t.Errorf("DeleteAt() val = %d, want %d", val, tt.wantVal)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.wantList) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.wantList)
			}
		})
	}
}

func TestSinglyLinkedList_Search(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		val     int
		want    int
	}{
		{
			name:    "search in empty list",
			initial: []int{},
			val:     1,
			want:    -1,
		},
		{
			name:    "search found at head",
			initial: []int{1, 2, 3},
			val:     1,
			want:    0,
		},
		{
			name:    "search found in middle",
			initial: []int{1, 2, 3},
			val:     2,
			want:    1,
		},
		{
			name:    "search found at tail",
			initial: []int{1, 2, 3},
			val:     3,
			want:    2,
		},
		{
			name:    "search not found",
			initial: []int{1, 2, 3},
			val:     4,
			want:    -1,
		},
		{
			name:    "search returns first occurrence",
			initial: []int{1, 2, 2, 3},
			val:     2,
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			got := l.Search(tt.val)
			if got != tt.want {
				t.Errorf("Search() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		want    []int
	}{
		{
			name:    "reverse empty list",
			initial: []int{},
			want:    []int{},
		},
		{
			name:    "reverse single element",
			initial: []int{1},
			want:    []int{1},
		},
		{
			name:    "reverse multiple elements",
			initial: []int{1, 2, 3, 4, 5},
			want:    []int{5, 4, 3, 2, 1},
		},
		{
			name:    "reverse two elements",
			initial: []int{1, 2},
			want:    []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			l.Reverse()

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() after Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ============================================================
// Doubly Linked List Tests
// ============================================================

func TestDoublyLinkedList_InsertAtHead(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "insert single element",
			values: []int{1},
			want:   []int{1},
		},
		{
			name:   "insert multiple elements",
			values: []int{1, 2, 3},
			want:   []int{3, 2, 1},
		},
		{
			name:   "insert no elements",
			values: []int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewDoublyLinkedList()
			for _, v := range tt.values {
				l.InsertAtHead(v)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
			if l.Len() != len(tt.want) {
				t.Errorf("Len() = %d, want %d", l.Len(), len(tt.want))
			}

			// Verify tail pointer
			if len(tt.values) > 0 && l.Tail.Val != tt.values[0] {
				t.Errorf("Tail.Val = %d, want %d", l.Tail.Val, tt.values[0])
			}
		})
	}
}

func TestDoublyLinkedList_InsertAtTail(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "insert single element",
			values: []int{1},
			want:   []int{1},
		},
		{
			name:   "insert multiple elements",
			values: []int{1, 2, 3},
			want:   []int{1, 2, 3},
		},
		{
			name:   "insert no elements",
			values: []int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewDoublyLinkedList()
			for _, v := range tt.values {
				l.InsertAtTail(v)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
			if l.Len() != len(tt.want) {
				t.Errorf("Len() = %d, want %d", l.Len(), len(tt.want))
			}

			// Verify head pointer
			if len(tt.values) > 0 && l.Head.Val != tt.values[0] {
				t.Errorf("Head.Val = %d, want %d", l.Head.Val, tt.values[0])
			}
		})
	}
}

func TestDoublyLinkedList_DeleteAtHead(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		wantVal  int
		wantErr  error
		wantList []int
	}{
		{
			name:     "delete from empty list",
			initial:  []int{},
			wantVal:  0,
			wantErr:  ErrEmptyList,
			wantList: []int{},
		},
		{
			name:     "delete single element",
			initial:  []int{1},
			wantVal:  1,
			wantErr:  nil,
			wantList: []int{},
		},
		{
			name:     "delete from multiple elements",
			initial:  []int{1, 2, 3},
			wantVal:  1,
			wantErr:  nil,
			wantList: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewDoublyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			val, err := l.DeleteAtHead()
			if err != tt.wantErr {
				t.Errorf("DeleteAtHead() error = %v, wantErr %v", err, tt.wantErr)
			}
			if val != tt.wantVal {
				t.Errorf("DeleteAtHead() val = %d, want %d", val, tt.wantVal)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.wantList) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.wantList)
			}

			// Verify head/tail consistency after deletion
			if l.Len() == 0 {
				if l.Head != nil {
					t.Errorf("Head should be nil for empty list")
				}
				if l.Tail != nil {
					t.Errorf("Tail should be nil for empty list")
				}
			}
			if l.Head != nil && l.Head.Prev != nil {
				t.Errorf("Head.Prev should be nil after DeleteAtHead")
			}
		})
	}
}

func TestDoublyLinkedList_DeleteAtTail(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		wantVal  int
		wantErr  error
		wantList []int
	}{
		{
			name:     "delete from empty list",
			initial:  []int{},
			wantVal:  0,
			wantErr:  ErrEmptyList,
			wantList: []int{},
		},
		{
			name:     "delete single element",
			initial:  []int{1},
			wantVal:  1,
			wantErr:  nil,
			wantList: []int{},
		},
		{
			name:     "delete from multiple elements",
			initial:  []int{1, 2, 3},
			wantVal:  3,
			wantErr:  nil,
			wantList: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewDoublyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			val, err := l.DeleteAtTail()
			if err != tt.wantErr {
				t.Errorf("DeleteAtTail() error = %v, wantErr %v", err, tt.wantErr)
			}
			if val != tt.wantVal {
				t.Errorf("DeleteAtTail() val = %d, want %d", val, tt.wantVal)
			}

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.wantList) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.wantList)
			}

			// Verify head/tail consistency after deletion
			if l.Len() == 0 {
				if l.Head != nil {
					t.Errorf("Head should be nil for empty list")
				}
				if l.Tail != nil {
					t.Errorf("Tail should be nil for empty list")
				}
			}
			if l.Tail != nil && l.Tail.Next != nil {
				t.Errorf("Tail.Next should be nil after DeleteAtTail")
			}
		})
	}
}

func TestDoublyLinkedList_Search(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		val     int
		want    int
	}{
		{
			name:    "search in empty list",
			initial: []int{},
			val:     1,
			want:    -1,
		},
		{
			name:    "search found at head",
			initial: []int{1, 2, 3},
			val:     1,
			want:    0,
		},
		{
			name:    "search found in middle",
			initial: []int{1, 2, 3},
			val:     2,
			want:    1,
		},
		{
			name:    "search found at tail",
			initial: []int{1, 2, 3},
			val:     3,
			want:    2,
		},
		{
			name:    "search not found",
			initial: []int{1, 2, 3},
			val:     4,
			want:    -1,
		},
		{
			name:    "search returns first occurrence",
			initial: []int{1, 2, 2, 3},
			val:     2,
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewDoublyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			got := l.Search(tt.val)
			if got != tt.want {
				t.Errorf("Search() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_Reverse(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		want    []int
	}{
		{
			name:    "reverse empty list",
			initial: []int{},
			want:    []int{},
		},
		{
			name:    "reverse single element",
			initial: []int{1},
			want:    []int{1},
		},
		{
			name:    "reverse multiple elements",
			initial: []int{1, 2, 3, 4, 5},
			want:    []int{5, 4, 3, 2, 1},
		},
		{
			name:    "reverse two elements",
			initial: []int{1, 2},
			want:    []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewDoublyLinkedList()
			for _, v := range tt.initial {
				l.InsertAtTail(v)
			}

			l.Reverse()

			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() after Reverse() = %v, want %v", got, tt.want)
			}

			// Verify head and tail pointers after reverse
			if len(tt.initial) > 0 {
				if l.Head.Val != tt.want[0] {
					t.Errorf("Head.Val = %d, want %d", l.Head.Val, tt.want[0])
				}
				if l.Tail.Val != tt.want[len(tt.want)-1] {
					t.Errorf("Tail.Val = %d, want %d", l.Tail.Val, tt.want[len(tt.want)-1])
				}
			}

			// Verify Prev pointers are correct after reverse
			if l.Head != nil && l.Head.Prev != nil {
				t.Errorf("Head.Prev should be nil after Reverse")
			}
			if l.Tail != nil && l.Tail.Next != nil {
				t.Errorf("Tail.Next should be nil after Reverse")
			}
		})
	}
}
