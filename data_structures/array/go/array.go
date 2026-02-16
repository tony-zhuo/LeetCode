package datastructures

import "errors"

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrEmptyArray      = errors.New("array is empty")
)

const defaultCapacity = 8

// DynamicArray is a manually managed dynamic array with explicit capacity control.
type DynamicArray struct {
	data []int
	size int
}

// NewDynamicArray creates a new empty dynamic array with default capacity.
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		data: make([]int, defaultCapacity),
		size: 0,
	}
}

// NewDynamicArrayWithCapacity creates a new empty dynamic array with specified initial capacity.
func NewDynamicArrayWithCapacity(cap int) *DynamicArray {
	if cap <= 0 {
		cap = defaultCapacity
	}
	return &DynamicArray{
		data: make([]int, cap),
		size: 0,
	}
}

// Get returns the element at the given index.
func (a *DynamicArray) Get(index int) (int, error) {
	if index < 0 || index >= a.size {
		return 0, ErrIndexOutOfRange
	}
	return a.data[index], nil
}

// Set updates the element at the given index.
func (a *DynamicArray) Set(index int, val int) error {
	if index < 0 || index >= a.size {
		return ErrIndexOutOfRange
	}
	a.data[index] = val
	return nil
}

// Append adds an element to the end of the array, resizing if necessary.
func (a *DynamicArray) Append(val int) {
	if a.size == len(a.data) {
		a.resize(len(a.data) * 2)
	}
	a.data[a.size] = val
	a.size++
}

// InsertAt inserts an element at the given index, shifting subsequent elements right.
func (a *DynamicArray) InsertAt(index int, val int) error {
	if index < 0 || index > a.size {
		return ErrIndexOutOfRange
	}
	if a.size == len(a.data) {
		a.resize(len(a.data) * 2)
	}
	copy(a.data[index+1:a.size+1], a.data[index:a.size])
	a.data[index] = val
	a.size++
	return nil
}

// DeleteAt removes the element at the given index, shifting subsequent elements left.
func (a *DynamicArray) DeleteAt(index int) (int, error) {
	if index < 0 || index >= a.size {
		return 0, ErrIndexOutOfRange
	}
	val := a.data[index]
	copy(a.data[index:a.size-1], a.data[index+1:a.size])
	a.size--
	// Shrink if size is 1/4 of capacity to avoid thrashing
	if a.size > 0 && a.size == len(a.data)/4 {
		a.resize(len(a.data) / 2)
	}
	return val, nil
}

// Size returns the number of elements in the array.
func (a *DynamicArray) Size() int {
	return a.size
}

// Capacity returns the current capacity of the underlying storage.
func (a *DynamicArray) Capacity() int {
	return len(a.data)
}

// IsEmpty returns true if the array contains no elements.
func (a *DynamicArray) IsEmpty() bool {
	return a.size == 0
}

// Contains returns true if the array contains the given value.
func (a *DynamicArray) Contains(val int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == val {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of val, or -1 if not found.
func (a *DynamicArray) IndexOf(val int) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == val {
			return i
		}
	}
	return -1
}

// Reverse reverses the array in place.
func (a *DynamicArray) Reverse() {
	for i, j := 0, a.size-1; i < j; i, j = i+1, j-1 {
		a.data[i], a.data[j] = a.data[j], a.data[i]
	}
}

// ToSlice returns a copy of the array as a slice.
func (a *DynamicArray) ToSlice() []int {
	result := make([]int, a.size)
	copy(result, a.data[:a.size])
	return result
}

func (a *DynamicArray) resize(newCap int) {
	newData := make([]int, newCap)
	copy(newData, a.data[:a.size])
	a.data = newData
}
