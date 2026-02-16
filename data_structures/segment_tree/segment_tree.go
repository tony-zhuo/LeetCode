package datastructures

import "math"

// SegmentTree supports range sum queries with point updates.
type SegmentTree struct {
	tree []int
	n    int
}

// NewSegmentTree builds a segment tree from the given slice of integers.
func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	st := &SegmentTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	if n > 0 {
		st.build(nums, 1, 0, n-1)
	}
	return st
}

// build recursively constructs the segment tree.
func (st *SegmentTree) build(nums []int, node, start, end int) {
	if start == end {
		st.tree[node] = nums[start]
		return
	}
	mid := (start + end) / 2
	st.build(nums, 2*node, start, mid)
	st.build(nums, 2*node+1, mid+1, end)
	st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
}

// Update sets the value at the given index and propagates the change up the tree.
func (st *SegmentTree) Update(index, val int) {
	st.update(1, 0, st.n-1, index, val)
}

// update recursively updates the segment tree for a point update.
func (st *SegmentTree) update(node, start, end, index, val int) {
	if start == end {
		st.tree[node] = val
		return
	}
	mid := (start + end) / 2
	if index <= mid {
		st.update(2*node, start, mid, index, val)
	} else {
		st.update(2*node+1, mid+1, end, index, val)
	}
	st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
}

// Query returns the sum of elements in the inclusive range [left, right].
func (st *SegmentTree) Query(left, right int) int {
	return st.query(1, 0, st.n-1, left, right)
}

// query recursively computes the range sum.
func (st *SegmentTree) query(node, start, end, left, right int) int {
	if right < start || end < left {
		return 0
	}
	if left <= start && end <= right {
		return st.tree[node]
	}
	mid := (start + end) / 2
	leftSum := st.query(2*node, start, mid, left, right)
	rightSum := st.query(2*node+1, mid+1, end, left, right)
	return leftSum + rightSum
}

// RangeMinSegmentTree supports range minimum queries with point updates.
type RangeMinSegmentTree struct {
	tree []int
	n    int
}

// NewRangeMinSegmentTree builds a range-minimum segment tree from the given slice.
func NewRangeMinSegmentTree(nums []int) *RangeMinSegmentTree {
	n := len(nums)
	st := &RangeMinSegmentTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	if n > 0 {
		for i := range st.tree {
			st.tree[i] = math.MaxInt
		}
		st.build(nums, 1, 0, n-1)
	}
	return st
}

// build recursively constructs the range-minimum segment tree.
func (st *RangeMinSegmentTree) build(nums []int, node, start, end int) {
	if start == end {
		st.tree[node] = nums[start]
		return
	}
	mid := (start + end) / 2
	st.build(nums, 2*node, start, mid)
	st.build(nums, 2*node+1, mid+1, end)
	st.tree[node] = min(st.tree[2*node], st.tree[2*node+1])
}

// Update sets the value at the given index and propagates the change up the tree.
func (st *RangeMinSegmentTree) Update(index, val int) {
	st.update(1, 0, st.n-1, index, val)
}

// update recursively updates the range-minimum segment tree for a point update.
func (st *RangeMinSegmentTree) update(node, start, end, index, val int) {
	if start == end {
		st.tree[node] = val
		return
	}
	mid := (start + end) / 2
	if index <= mid {
		st.update(2*node, start, mid, index, val)
	} else {
		st.update(2*node+1, mid+1, end, index, val)
	}
	st.tree[node] = min(st.tree[2*node], st.tree[2*node+1])
}

// Query returns the minimum element in the inclusive range [left, right].
func (st *RangeMinSegmentTree) Query(left, right int) int {
	return st.query(1, 0, st.n-1, left, right)
}

// query recursively computes the range minimum.
func (st *RangeMinSegmentTree) query(node, start, end, left, right int) int {
	if right < start || end < left {
		return math.MaxInt
	}
	if left <= start && end <= right {
		return st.tree[node]
	}
	mid := (start + end) / 2
	leftMin := st.query(2*node, start, mid, left, right)
	rightMin := st.query(2*node+1, mid+1, end, left, right)
	return min(leftMin, rightMin)
}
