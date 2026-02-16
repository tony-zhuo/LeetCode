package datastructures

// UnionFind is a disjoint set data structure that tracks elements partitioned
// into non-overlapping subsets. It supports near-constant time union and find
// operations using path compression and union by rank optimizations.
type UnionFind struct {
	parent []int
	rank   []int
	count  int // number of disjoint sets
}

// NewUnionFind creates a new UnionFind with n elements (0 to n-1).
// Each element starts as its own parent with rank 0, and count equals n.
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
		count:  n,
	}
}

// Find returns the root representative of the set containing x.
// It uses path compression to set the parent of each traversed node directly
// to the root, flattening the tree for future lookups.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y using union by rank.
// It attaches the shorter tree under the root of the taller tree.
// Returns false if x and y are already in the same set; true otherwise.
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	uf.count--
	return true
}

// Connected returns true if x and y belong to the same set.
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Count returns the current number of disjoint sets.
func (uf *UnionFind) Count() int {
	return uf.count
}
