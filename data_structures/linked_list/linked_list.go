package datastructures

import (
	"errors"
)

// ============================================================
// ListNode — simple linked list node used by problem solutions
// ============================================================

// ListNode is a simple linked list node used by problem solutions.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Arr2Node converts a slice of ints into a linked list and returns the head.
func Arr2Node(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  input[0],
		Next: nil,
	}
	curr := head

	for i := 1; i < len(input); i++ {
		curr.Next = &ListNode{
			Val:  input[i],
			Next: nil,
		}

		curr = curr.Next
	}

	return head
}

var (
	ErrEmptyList      = errors.New("list is empty")
	ErrIndexOutOfRange = errors.New("index out of range")
)

// ============================================================
// Singly Linked List
// ============================================================

type SinglyNode struct {
	Val  int
	Next *SinglyNode
}

type SinglyLinkedList struct {
	Head *SinglyNode
	size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (l *SinglyLinkedList) InsertAtHead(val int) {
	node := &SinglyNode{Val: val, Next: l.Head}
	l.Head = node
	l.size++
}

func (l *SinglyLinkedList) InsertAtTail(val int) {
	node := &SinglyNode{Val: val}

	if l.Head == nil {
		l.Head = node
		l.size++
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	l.size++
}

func (l *SinglyLinkedList) InsertAt(index, val int) error {
	if index < 0 || index > l.size {
		return ErrIndexOutOfRange
	}

	if index == 0 {
		l.InsertAtHead(val)
		return nil
	}

	curr := l.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}

	node := &SinglyNode{Val: val, Next: curr.Next}
	curr.Next = node
	l.size++
	return nil
}

func (l *SinglyLinkedList) DeleteAtHead() (int, error) {
	if l.Head == nil {
		return 0, ErrEmptyList
	}

	val := l.Head.Val
	l.Head = l.Head.Next
	l.size--
	return val, nil
}

func (l *SinglyLinkedList) DeleteAtTail() (int, error) {
	if l.Head == nil {
		return 0, ErrEmptyList
	}

	if l.Head.Next == nil {
		val := l.Head.Val
		l.Head = nil
		l.size--
		return val, nil
	}

	curr := l.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}

	val := curr.Next.Val
	curr.Next = nil
	l.size--
	return val, nil
}

func (l *SinglyLinkedList) DeleteAt(index int) (int, error) {
	if l.Head == nil {
		return 0, ErrEmptyList
	}

	if index < 0 || index >= l.size {
		return 0, ErrIndexOutOfRange
	}

	if index == 0 {
		return l.DeleteAtHead()
	}

	curr := l.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}

	val := curr.Next.Val
	curr.Next = curr.Next.Next
	l.size--
	return val, nil
}

func (l *SinglyLinkedList) Search(val int) int {
	curr := l.Head
	for i := 0; curr != nil; i++ {
		if curr.Val == val {
			return i
		}
		curr = curr.Next
	}
	return -1
}

func (l *SinglyLinkedList) Reverse() {
	var prev *SinglyNode
	curr := l.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l.Head = prev
}

func (l *SinglyLinkedList) Len() int {
	return l.size
}

func (l *SinglyLinkedList) ToSlice() []int {
	result := make([]int, 0, l.size)
	curr := l.Head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

// ============================================================
// Doubly Linked List
// ============================================================

type DoublyNode struct {
	Val  int
	Prev *DoublyNode
	Next *DoublyNode
}

type DoublyLinkedList struct {
	Head *DoublyNode
	Tail *DoublyNode
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) InsertAtHead(val int) {
	node := &DoublyNode{Val: val, Next: l.Head}

	if l.Head != nil {
		l.Head.Prev = node
	} else {
		l.Tail = node
	}

	l.Head = node
	l.size++
}

func (l *DoublyLinkedList) InsertAtTail(val int) {
	node := &DoublyNode{Val: val, Prev: l.Tail}

	if l.Tail != nil {
		l.Tail.Next = node
	} else {
		l.Head = node
	}

	l.Tail = node
	l.size++
}

func (l *DoublyLinkedList) DeleteAtHead() (int, error) {
	if l.Head == nil {
		return 0, ErrEmptyList
	}

	val := l.Head.Val
	l.Head = l.Head.Next

	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}

	l.size--
	return val, nil
}

func (l *DoublyLinkedList) DeleteAtTail() (int, error) {
	if l.Tail == nil {
		return 0, ErrEmptyList
	}

	val := l.Tail.Val
	l.Tail = l.Tail.Prev

	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}

	l.size--
	return val, nil
}

func (l *DoublyLinkedList) Search(val int) int {
	curr := l.Head
	for i := 0; curr != nil; i++ {
		if curr.Val == val {
			return i
		}
		curr = curr.Next
	}
	return -1
}

func (l *DoublyLinkedList) Reverse() {
	curr := l.Head

	for curr != nil {
		curr.Prev, curr.Next = curr.Next, curr.Prev
		curr = curr.Prev
	}

	l.Head, l.Tail = l.Tail, l.Head
}

func (l *DoublyLinkedList) Len() int {
	return l.size
}

func (l *DoublyLinkedList) ToSlice() []int {
	result := make([]int, 0, l.size)
	curr := l.Head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}
