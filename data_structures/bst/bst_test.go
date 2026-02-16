package datastructures

import (
	"errors"
	"reflect"
	"testing"
)

func TestBST_InsertAndInorderTraversal(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		want       []int
	}{
		{
			name:       "single element",
			insertVals: []int{5},
			want:       []int{5},
		},
		{
			name:       "already sorted input",
			insertVals: []int{1, 2, 3, 4, 5},
			want:       []int{1, 2, 3, 4, 5},
		},
		{
			name:       "reverse sorted input",
			insertVals: []int{5, 4, 3, 2, 1},
			want:       []int{1, 2, 3, 4, 5},
		},
		{
			name:       "random order input",
			insertVals: []int{5, 3, 7, 1, 4, 6, 8},
			want:       []int{1, 3, 4, 5, 6, 7, 8},
		},
		{
			name:       "negative values",
			insertVals: []int{0, -3, 5, -1, 2},
			want:       []int{-3, -1, 0, 2, 5},
		},
		{
			name:       "empty tree",
			insertVals: []int{},
			want:       []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.insertVals {
				bst.Insert(v)
			}

			got := bst.InorderTraversal()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_InsertDuplicateIgnored(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		wantOrder  []int
		wantSize   int
	}{
		{
			name:       "duplicate at root",
			insertVals: []int{5, 5},
			wantOrder:  []int{5},
			wantSize:   1,
		},
		{
			name:       "multiple duplicates mixed",
			insertVals: []int{5, 3, 7, 3, 5, 7, 1},
			wantOrder:  []int{1, 3, 5, 7},
			wantSize:   4,
		},
		{
			name:       "all same values",
			insertVals: []int{4, 4, 4, 4},
			wantOrder:  []int{4},
			wantSize:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.insertVals {
				bst.Insert(v)
			}

			got := bst.InorderTraversal()
			if !reflect.DeepEqual(got, tt.wantOrder) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.wantOrder)
			}
			if bst.Size() != tt.wantSize {
				t.Errorf("Size() = %d, want %d", bst.Size(), tt.wantSize)
			}
		})
	}
}

func TestBST_Search(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		searchVal  int
		want       bool
	}{
		{
			name:       "find existing root",
			insertVals: []int{5, 3, 7},
			searchVal:  5,
			want:       true,
		},
		{
			name:       "find existing left child",
			insertVals: []int{5, 3, 7},
			searchVal:  3,
			want:       true,
		},
		{
			name:       "find existing right child",
			insertVals: []int{5, 3, 7},
			searchVal:  7,
			want:       true,
		},
		{
			name:       "value not found",
			insertVals: []int{5, 3, 7},
			searchVal:  10,
			want:       false,
		},
		{
			name:       "search in empty tree",
			insertVals: []int{},
			searchVal:  1,
			want:       false,
		},
		{
			name:       "find deep node",
			insertVals: []int{10, 5, 15, 3, 7, 12, 20, 1},
			searchVal:  1,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.insertVals {
				bst.Insert(v)
			}

			if got := bst.Search(tt.searchVal); got != tt.want {
				t.Errorf("Search(%d) = %v, want %v", tt.searchVal, got, tt.want)
			}
		})
	}
}

func TestBST_DeleteLeaf(t *testing.T) {
	tests := []struct {
		name      string
		insertVals []int
		deleteVal int
		wantOrder []int
		wantSize  int
	}{
		{
			name:       "delete left leaf",
			insertVals: []int{5, 3, 7},
			deleteVal:  3,
			wantOrder:  []int{5, 7},
			wantSize:   2,
		},
		{
			name:       "delete right leaf",
			insertVals: []int{5, 3, 7},
			deleteVal:  7,
			wantOrder:  []int{3, 5},
			wantSize:   2,
		},
		{
			name:       "delete only node (root leaf)",
			insertVals: []int{5},
			deleteVal:  5,
			wantOrder:  []int{},
			wantSize:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.insertVals {
				bst.Insert(v)
			}

			deleted := bst.Delete(tt.deleteVal)
			if !deleted {
				t.Fatalf("Delete(%d) = false, want true", tt.deleteVal)
			}

			got := bst.InorderTraversal()
			if !reflect.DeepEqual(got, tt.wantOrder) {
				t.Errorf("InorderTraversal() after delete = %v, want %v", got, tt.wantOrder)
			}
			if bst.Size() != tt.wantSize {
				t.Errorf("Size() after delete = %d, want %d", bst.Size(), tt.wantSize)
			}
		})
	}
}

func TestBST_DeleteOneChild(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		deleteVal  int
		wantOrder  []int
		wantSize   int
	}{
		{
			name:       "delete node with only left child",
			insertVals: []int{10, 5, 3},
			deleteVal:  5,
			wantOrder:  []int{3, 10},
			wantSize:   2,
		},
		{
			name:       "delete node with only right child",
			insertVals: []int{10, 5, 7},
			deleteVal:  5,
			wantOrder:  []int{7, 10},
			wantSize:   2,
		},
		{
			name:       "delete root with only left child",
			insertVals: []int{10, 5, 3},
			deleteVal:  10,
			wantOrder:  []int{3, 5},
			wantSize:   2,
		},
		{
			name:       "delete root with only right child",
			insertVals: []int{10, 15, 20},
			deleteVal:  10,
			wantOrder:  []int{15, 20},
			wantSize:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.insertVals {
				bst.Insert(v)
			}

			deleted := bst.Delete(tt.deleteVal)
			if !deleted {
				t.Fatalf("Delete(%d) = false, want true", tt.deleteVal)
			}

			got := bst.InorderTraversal()
			if !reflect.DeepEqual(got, tt.wantOrder) {
				t.Errorf("InorderTraversal() after delete = %v, want %v", got, tt.wantOrder)
			}
			if bst.Size() != tt.wantSize {
				t.Errorf("Size() after delete = %d, want %d", bst.Size(), tt.wantSize)
			}
		})
	}
}

func TestBST_DeleteTwoChildren(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		deleteVal  int
		wantOrder  []int
		wantSize   int
	}{
		{
			name:       "delete node with two children",
			insertVals: []int{10, 5, 15, 3, 7},
			deleteVal:  5,
			wantOrder:  []int{3, 7, 10, 15},
			wantSize:   4,
		},
		{
			name:       "delete root with two children",
			insertVals: []int{10, 5, 15, 3, 7, 12, 20},
			deleteVal:  10,
			wantOrder:  []int{3, 5, 7, 12, 15, 20},
			wantSize:   6,
		},
		{
			name:       "delete node where successor is immediate right child",
			insertVals: []int{10, 5, 15, 3, 12, 20},
			deleteVal:  15,
			wantOrder:  []int{3, 5, 10, 12, 20},
			wantSize:   5,
		},
		{
			name:       "delete node where successor has a right subtree",
			insertVals: []int{20, 10, 30, 5, 15, 25, 35, 22, 27},
			deleteVal:  20,
			wantOrder:  []int{5, 10, 15, 22, 25, 27, 30, 35},
			wantSize:   8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.insertVals {
				bst.Insert(v)
			}

			deleted := bst.Delete(tt.deleteVal)
			if !deleted {
				t.Fatalf("Delete(%d) = false, want true", tt.deleteVal)
			}

			got := bst.InorderTraversal()
			if !reflect.DeepEqual(got, tt.wantOrder) {
				t.Errorf("InorderTraversal() after delete = %v, want %v", got, tt.wantOrder)
			}
			if bst.Size() != tt.wantSize {
				t.Errorf("Size() after delete = %d, want %d", bst.Size(), tt.wantSize)
			}
		})
	}
}

func TestBST_DeleteNotFound(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		deleteVal  int
	}{
		{
			name:       "delete from empty tree",
			insertVals: []int{},
			deleteVal:  5,
		},
		{
			name:       "delete non-existing value",
			insertVals: []int{10, 5, 15},
			deleteVal:  7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.insertVals {
				bst.Insert(v)
			}

			origSize := bst.Size()
			deleted := bst.Delete(tt.deleteVal)
			if deleted {
				t.Errorf("Delete(%d) = true, want false", tt.deleteVal)
			}
			if bst.Size() != origSize {
				t.Errorf("Size() changed after failed delete: got %d, want %d", bst.Size(), origSize)
			}
		})
	}
}

func TestBST_MinMax_EmptyTree(t *testing.T) {
	tests := []struct {
		name string
		fn   func(b *BST) (int, error)
	}{
		{
			name: "Min on empty tree",
			fn:   func(b *BST) (int, error) { return b.Min() },
		},
		{
			name: "Max on empty tree",
			fn:   func(b *BST) (int, error) { return b.Max() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			_, err := tt.fn(bst)
			if !errors.Is(err, ErrEmptyTree) {
				t.Errorf("error = %v, want %v", err, ErrEmptyTree)
			}
		})
	}
}

func TestBST_MinMax_PopulatedTree(t *testing.T) {
	tests := []struct {
		name       string
		insertVals []int
		wantMin    int
		wantMax    int
	}{
		{
			name:       "single element",
			insertVals: []int{5},
			wantMin:    5,
			wantMax:    5,
		},
		{
			name:       "balanced tree",
			insertVals: []int{10, 5, 15, 3, 7, 12, 20},
			wantMin:    3,
			wantMax:    20,
		},
		{
			name:       "left-skewed tree",
			insertVals: []int{5, 4, 3, 2, 1},
			wantMin:    1,
			wantMax:    5,
		},
		{
			name:       "right-skewed tree",
			insertVals: []int{1, 2, 3, 4, 5},
			wantMin:    1,
			wantMax:    5,
		},
		{
			name:       "negative values",
			insertVals: []int{0, -10, 10, -20, -5},
			wantMin:    -20,
			wantMax:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.insertVals {
				bst.Insert(v)
			}

			gotMin, err := bst.Min()
			if err != nil {
				t.Fatalf("Min() unexpected error: %v", err)
			}
			if gotMin != tt.wantMin {
				t.Errorf("Min() = %d, want %d", gotMin, tt.wantMin)
			}

			gotMax, err := bst.Max()
			if err != nil {
				t.Fatalf("Max() unexpected error: %v", err)
			}
			if gotMax != tt.wantMax {
				t.Errorf("Max() = %d, want %d", gotMax, tt.wantMax)
			}
		})
	}
}

func TestBST_Size(t *testing.T) {
	tests := []struct {
		name       string
		operations func(b *BST)
		wantSize   int
	}{
		{
			name:       "empty tree has size 0",
			operations: func(b *BST) {},
			wantSize:   0,
		},
		{
			name: "size after inserts",
			operations: func(b *BST) {
				b.Insert(5)
				b.Insert(3)
				b.Insert(7)
			},
			wantSize: 3,
		},
		{
			name: "size unchanged after duplicate insert",
			operations: func(b *BST) {
				b.Insert(5)
				b.Insert(5)
				b.Insert(5)
			},
			wantSize: 1,
		},
		{
			name: "size decreases after delete",
			operations: func(b *BST) {
				b.Insert(5)
				b.Insert(3)
				b.Insert(7)
				b.Delete(3)
			},
			wantSize: 2,
		},
		{
			name: "size unchanged after failed delete",
			operations: func(b *BST) {
				b.Insert(5)
				b.Insert(3)
				b.Delete(10)
			},
			wantSize: 2,
		},
		{
			name: "size returns to 0 after deleting all",
			operations: func(b *BST) {
				b.Insert(5)
				b.Insert(3)
				b.Insert(7)
				b.Delete(3)
				b.Delete(7)
				b.Delete(5)
			},
			wantSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			tt.operations(bst)
			if got := bst.Size(); got != tt.wantSize {
				t.Errorf("Size() = %d, want %d", got, tt.wantSize)
			}
		})
	}
}
