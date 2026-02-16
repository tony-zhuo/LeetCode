package datastructures

import (
	"errors"
	"reflect"
	"testing"
)

func TestDynamicArray_AppendAndGet(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		getIndex int
		want     int
	}{
		{
			name:     "single element",
			values:   []int{10},
			getIndex: 0,
			want:     10,
		},
		{
			name:     "multiple elements get first",
			values:   []int{1, 2, 3, 4, 5},
			getIndex: 0,
			want:     1,
		},
		{
			name:     "multiple elements get last",
			values:   []int{1, 2, 3, 4, 5},
			getIndex: 4,
			want:     5,
		},
		{
			name:     "multiple elements get middle",
			values:   []int{10, 20, 30},
			getIndex: 1,
			want:     20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewDynamicArray()
			for _, v := range tt.values {
				a.Append(v)
			}
			got, err := a.Get(tt.getIndex)
			if err != nil {
				t.Fatalf("Get(%d) unexpected error: %v", tt.getIndex, err)
			}
			if got != tt.want {
				t.Errorf("Get(%d) = %d, want %d", tt.getIndex, got, tt.want)
			}
		})
	}
}

func TestDynamicArray_GetOutOfRange(t *testing.T) {
	tests := []struct {
		name  string
		size  int
		index int
	}{
		{name: "negative index", size: 3, index: -1},
		{name: "index equals size", size: 3, index: 3},
		{name: "index beyond size", size: 3, index: 10},
		{name: "empty array index 0", size: 0, index: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewDynamicArray()
			for i := 0; i < tt.size; i++ {
				a.Append(i)
			}
			_, err := a.Get(tt.index)
			if !errors.Is(err, ErrIndexOutOfRange) {
				t.Errorf("Get(%d) error = %v, want %v", tt.index, err, ErrIndexOutOfRange)
			}
		})
	}
}

func TestDynamicArray_Set(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		setIndex int
		setVal   int
		wantArr  []int
	}{
		{
			name:     "set first element",
			initial:  []int{1, 2, 3},
			setIndex: 0,
			setVal:   99,
			wantArr:  []int{99, 2, 3},
		},
		{
			name:     "set last element",
			initial:  []int{1, 2, 3},
			setIndex: 2,
			setVal:   99,
			wantArr:  []int{1, 2, 99},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewDynamicArray()
			for _, v := range tt.initial {
				a.Append(v)
			}
			if err := a.Set(tt.setIndex, tt.setVal); err != nil {
				t.Fatalf("Set(%d, %d) unexpected error: %v", tt.setIndex, tt.setVal, err)
			}
			if got := a.ToSlice(); !reflect.DeepEqual(got, tt.wantArr) {
				t.Errorf("after Set: got %v, want %v", got, tt.wantArr)
			}
		})
	}
}

func TestDynamicArray_InsertAt(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		index   int
		val     int
		wantArr []int
	}{
		{
			name:    "insert at beginning",
			initial: []int{2, 3, 4},
			index:   0,
			val:     1,
			wantArr: []int{1, 2, 3, 4},
		},
		{
			name:    "insert at middle",
			initial: []int{1, 3, 4},
			index:   1,
			val:     2,
			wantArr: []int{1, 2, 3, 4},
		},
		{
			name:    "insert at end",
			initial: []int{1, 2, 3},
			index:   3,
			val:     4,
			wantArr: []int{1, 2, 3, 4},
		},
		{
			name:    "insert into empty array",
			initial: []int{},
			index:   0,
			val:     1,
			wantArr: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewDynamicArray()
			for _, v := range tt.initial {
				a.Append(v)
			}
			if err := a.InsertAt(tt.index, tt.val); err != nil {
				t.Fatalf("InsertAt(%d, %d) unexpected error: %v", tt.index, tt.val, err)
			}
			if got := a.ToSlice(); !reflect.DeepEqual(got, tt.wantArr) {
				t.Errorf("after InsertAt: got %v, want %v", got, tt.wantArr)
			}
		})
	}
}

func TestDynamicArray_InsertAt_OutOfRange(t *testing.T) {
	a := NewDynamicArray()
	a.Append(1)
	a.Append(2)

	if err := a.InsertAt(-1, 0); !errors.Is(err, ErrIndexOutOfRange) {
		t.Errorf("InsertAt(-1) error = %v, want %v", err, ErrIndexOutOfRange)
	}
	if err := a.InsertAt(3, 0); !errors.Is(err, ErrIndexOutOfRange) {
		t.Errorf("InsertAt(3) error = %v, want %v", err, ErrIndexOutOfRange)
	}
}

func TestDynamicArray_DeleteAt(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		index   int
		wantVal int
		wantArr []int
	}{
		{
			name:    "delete first",
			initial: []int{1, 2, 3},
			index:   0,
			wantVal: 1,
			wantArr: []int{2, 3},
		},
		{
			name:    "delete middle",
			initial: []int{1, 2, 3},
			index:   1,
			wantVal: 2,
			wantArr: []int{1, 3},
		},
		{
			name:    "delete last",
			initial: []int{1, 2, 3},
			index:   2,
			wantVal: 3,
			wantArr: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewDynamicArray()
			for _, v := range tt.initial {
				a.Append(v)
			}
			got, err := a.DeleteAt(tt.index)
			if err != nil {
				t.Fatalf("DeleteAt(%d) unexpected error: %v", tt.index, err)
			}
			if got != tt.wantVal {
				t.Errorf("DeleteAt(%d) = %d, want %d", tt.index, got, tt.wantVal)
			}
			if arr := a.ToSlice(); !reflect.DeepEqual(arr, tt.wantArr) {
				t.Errorf("after DeleteAt: got %v, want %v", arr, tt.wantArr)
			}
		})
	}
}

func TestDynamicArray_Resize(t *testing.T) {
	a := NewDynamicArrayWithCapacity(2)
	if a.Capacity() != 2 {
		t.Fatalf("initial Capacity() = %d, want 2", a.Capacity())
	}

	// Fill to trigger grow
	a.Append(1)
	a.Append(2)
	a.Append(3) // triggers resize to 4

	if a.Capacity() != 4 {
		t.Errorf("after grow Capacity() = %d, want 4", a.Capacity())
	}
	if got := a.ToSlice(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("after grow ToSlice() = %v, want [1 2 3]", got)
	}
}

func TestDynamicArray_Contains(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		search int
		want   bool
	}{
		{name: "found at beginning", values: []int{1, 2, 3}, search: 1, want: true},
		{name: "found at end", values: []int{1, 2, 3}, search: 3, want: true},
		{name: "not found", values: []int{1, 2, 3}, search: 4, want: false},
		{name: "empty array", values: []int{}, search: 1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewDynamicArray()
			for _, v := range tt.values {
				a.Append(v)
			}
			if got := a.Contains(tt.search); got != tt.want {
				t.Errorf("Contains(%d) = %v, want %v", tt.search, got, tt.want)
			}
		})
	}
}

func TestDynamicArray_IndexOf(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		search int
		want   int
	}{
		{name: "found first", values: []int{10, 20, 30}, search: 10, want: 0},
		{name: "found last", values: []int{10, 20, 30}, search: 30, want: 2},
		{name: "not found", values: []int{10, 20, 30}, search: 99, want: -1},
		{name: "duplicate returns first", values: []int{1, 2, 1, 3}, search: 1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewDynamicArray()
			for _, v := range tt.values {
				a.Append(v)
			}
			if got := a.IndexOf(tt.search); got != tt.want {
				t.Errorf("IndexOf(%d) = %d, want %d", tt.search, got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Reverse(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{name: "odd length", values: []int{1, 2, 3}, want: []int{3, 2, 1}},
		{name: "even length", values: []int{1, 2, 3, 4}, want: []int{4, 3, 2, 1}},
		{name: "single element", values: []int{1}, want: []int{1}},
		{name: "empty", values: []int{}, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewDynamicArray()
			for _, v := range tt.values {
				a.Append(v)
			}
			a.Reverse()
			if got := a.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("after Reverse: got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_SizeAndIsEmpty(t *testing.T) {
	a := NewDynamicArray()
	if !a.IsEmpty() {
		t.Error("new array should be empty")
	}
	if a.Size() != 0 {
		t.Errorf("new array Size() = %d, want 0", a.Size())
	}

	a.Append(1)
	a.Append(2)
	if a.IsEmpty() {
		t.Error("array with elements should not be empty")
	}
	if a.Size() != 2 {
		t.Errorf("Size() = %d, want 2", a.Size())
	}
}
