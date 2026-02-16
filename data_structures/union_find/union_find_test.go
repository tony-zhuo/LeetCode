package datastructures

import "testing"

func TestNewUnionFind(t *testing.T) {
	t.Run("count equals n", func(t *testing.T) {
		uf := NewUnionFind(5)
		if uf.Count() != 5 {
			t.Errorf("expected count 5, got %d", uf.Count())
		}
	})

	t.Run("each element is its own set", func(t *testing.T) {
		uf := NewUnionFind(5)
		for i := 0; i < 5; i++ {
			if uf.Find(i) != i {
				t.Errorf("expected Find(%d) = %d, got %d", i, i, uf.Find(i))
			}
		}
	})

	t.Run("no two distinct elements are connected initially", func(t *testing.T) {
		uf := NewUnionFind(5)
		for i := 0; i < 5; i++ {
			for j := i + 1; j < 5; j++ {
				if uf.Connected(i, j) {
					t.Errorf("expected %d and %d to be disconnected", i, j)
				}
			}
		}
	})
}

func TestUnion(t *testing.T) {
	t.Run("union two elements decreases count by 1", func(t *testing.T) {
		uf := NewUnionFind(5)
		result := uf.Union(0, 1)
		if !result {
			t.Error("expected Union(0, 1) to return true")
		}
		if uf.Count() != 4 {
			t.Errorf("expected count 4, got %d", uf.Count())
		}
	})

	t.Run("union two elements makes them connected", func(t *testing.T) {
		uf := NewUnionFind(5)
		uf.Union(0, 1)
		if !uf.Connected(0, 1) {
			t.Error("expected 0 and 1 to be connected after union")
		}
	})

	t.Run("union already connected elements returns false", func(t *testing.T) {
		uf := NewUnionFind(5)
		uf.Union(0, 1)
		result := uf.Union(0, 1)
		if result {
			t.Error("expected Union(0, 1) to return false when already connected")
		}
	})

	t.Run("union already connected elements does not change count", func(t *testing.T) {
		uf := NewUnionFind(5)
		uf.Union(0, 1)
		countBefore := uf.Count()
		uf.Union(0, 1)
		if uf.Count() != countBefore {
			t.Errorf("expected count %d, got %d", countBefore, uf.Count())
		}
	})
}

func TestFindWithPathCompression(t *testing.T) {
	t.Run("path compression updates parent to root", func(t *testing.T) {
		uf := NewUnionFind(5)
		// Build a chain: 0 <- 1 <- 2 <- 3
		uf.Union(0, 1)
		uf.Union(1, 2)
		uf.Union(2, 3)

		root := uf.Find(3)
		// After Find(3) with path compression, parent of 3 should point directly to root
		if uf.parent[3] != root {
			t.Errorf("expected parent[3] = %d (root), got %d", root, uf.parent[3])
		}
	})

	t.Run("all nodes in chain point to root after find", func(t *testing.T) {
		uf := NewUnionFind(6)
		// Create chain: union(0,1), union(1,2), union(2,3), union(3,4), union(4,5)
		uf.Union(0, 1)
		uf.Union(1, 2)
		uf.Union(2, 3)
		uf.Union(3, 4)
		uf.Union(4, 5)

		root := uf.Find(5)
		// After finding root from 5, verify path compression occurred
		for i := 0; i <= 5; i++ {
			// Call Find to trigger path compression on each node
			uf.Find(i)
			if uf.parent[i] != root {
				t.Errorf("expected parent[%d] = %d (root) after path compression, got %d", i, root, uf.parent[i])
			}
		}
	})
}

func TestTransitivity(t *testing.T) {
	t.Run("if a-b and b-c then a-c connected", func(t *testing.T) {
		uf := NewUnionFind(5)
		uf.Union(0, 1)
		uf.Union(1, 2)
		if !uf.Connected(0, 2) {
			t.Error("expected 0 and 2 to be connected via transitivity")
		}
	})

	t.Run("longer chain transitivity", func(t *testing.T) {
		uf := NewUnionFind(6)
		uf.Union(0, 1)
		uf.Union(1, 2)
		uf.Union(2, 3)
		uf.Union(3, 4)
		uf.Union(4, 5)

		// All elements should be connected
		for i := 0; i < 6; i++ {
			for j := i + 1; j < 6; j++ {
				if !uf.Connected(i, j) {
					t.Errorf("expected %d and %d to be connected", i, j)
				}
			}
		}
		if uf.Count() != 1 {
			t.Errorf("expected count 1, got %d", uf.Count())
		}
	})
}

func TestLargeUnionFindWithMultipleComponents(t *testing.T) {
	t.Run("multiple disjoint components", func(t *testing.T) {
		n := 100
		uf := NewUnionFind(n)

		// Create 10 components of size 10: {0-9}, {10-19}, ..., {90-99}
		for start := 0; start < n; start += 10 {
			for i := start + 1; i < start+10; i++ {
				uf.Union(start, i)
			}
		}

		if uf.Count() != 10 {
			t.Errorf("expected 10 components, got %d", uf.Count())
		}

		// Elements within the same component should be connected
		for start := 0; start < n; start += 10 {
			for i := start; i < start+10; i++ {
				for j := i + 1; j < start+10; j++ {
					if !uf.Connected(i, j) {
						t.Errorf("expected %d and %d to be connected within component", i, j)
					}
				}
			}
		}

		// Elements in different components should not be connected
		if uf.Connected(0, 10) {
			t.Error("expected 0 and 10 to be in different components")
		}
		if uf.Connected(5, 15) {
			t.Error("expected 5 and 15 to be in different components")
		}
		if uf.Connected(9, 90) {
			t.Error("expected 9 and 90 to be in different components")
		}
	})

	t.Run("merging components reduces count correctly", func(t *testing.T) {
		n := 100
		uf := NewUnionFind(n)

		// Create 10 components of size 10
		for start := 0; start < n; start += 10 {
			for i := start + 1; i < start+10; i++ {
				uf.Union(start, i)
			}
		}

		// Merge first two components
		uf.Union(0, 10)
		if uf.Count() != 9 {
			t.Errorf("expected 9 components after merging two, got %d", uf.Count())
		}

		// Verify cross-component connectivity
		if !uf.Connected(5, 15) {
			t.Error("expected 5 and 15 to be connected after merging their components")
		}

		// Merge all components into one
		for start := 20; start < n; start += 10 {
			uf.Union(0, start)
		}
		if uf.Count() != 1 {
			t.Errorf("expected 1 component after merging all, got %d", uf.Count())
		}
	})
}
