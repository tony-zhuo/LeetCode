package datastructures

import (
	"reflect"
	"sort"
	"testing"
)

func TestUndirectedGraph_AddEdge(t *testing.T) {
	t.Run("AddEdge creates edges in both directions", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge(1, 2)

		if !g.HasEdge(1, 2) {
			t.Error("expected edge from 1 to 2")
		}
		if !g.HasEdge(2, 1) {
			t.Error("expected edge from 2 to 1 (undirected)")
		}
	})
}

func TestDirectedGraph_AddEdge(t *testing.T) {
	t.Run("AddEdge creates edge in one direction only", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge(1, 2)

		if !g.HasEdge(1, 2) {
			t.Error("expected edge from 1 to 2")
		}
		if g.HasEdge(2, 1) {
			t.Error("did not expect edge from 2 to 1 (directed)")
		}
	})
}

func TestGraph_AddVertex(t *testing.T) {
	t.Run("AddVertex adds isolated vertex", func(t *testing.T) {
		g := NewGraph(false)
		g.AddVertex(5)

		vertices := g.Vertices()
		if !reflect.DeepEqual(vertices, []int{5}) {
			t.Errorf("expected vertices [5], got %v", vertices)
		}
	})

	t.Run("AddVertex does not overwrite existing neighbors", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge(1, 2)
		g.AddVertex(1)

		if !g.HasEdge(1, 2) {
			t.Error("AddVertex should not overwrite existing adjacency list")
		}
	})
}

func TestGraph_Vertices(t *testing.T) {
	t.Run("Vertices returns sorted list", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge(3, 1)
		g.AddEdge(2, 4)

		vertices := g.Vertices()
		expected := []int{1, 2, 3, 4}
		if !reflect.DeepEqual(vertices, expected) {
			t.Errorf("expected %v, got %v", expected, vertices)
		}
	})
}

func TestGraph_RemoveEdge(t *testing.T) {
	t.Run("RemoveEdge in undirected graph removes both directions", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge(1, 2)
		g.RemoveEdge(1, 2)

		if g.HasEdge(1, 2) {
			t.Error("edge from 1 to 2 should be removed")
		}
		if g.HasEdge(2, 1) {
			t.Error("edge from 2 to 1 should be removed (undirected)")
		}
	})

	t.Run("RemoveEdge in directed graph removes one direction", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge(1, 2)
		g.AddEdge(2, 1)
		g.RemoveEdge(1, 2)

		if g.HasEdge(1, 2) {
			t.Error("edge from 1 to 2 should be removed")
		}
		if !g.HasEdge(2, 1) {
			t.Error("edge from 2 to 1 should still exist (directed)")
		}
	})
}

func TestGraph_HasEdge(t *testing.T) {
	t.Run("HasEdge returns false for non-existent vertex", func(t *testing.T) {
		g := NewGraph(false)
		if g.HasEdge(1, 2) {
			t.Error("expected false for non-existent vertices")
		}
	})

	t.Run("HasEdge returns false for non-existent edge", func(t *testing.T) {
		g := NewGraph(false)
		g.AddVertex(1)
		g.AddVertex(2)
		if g.HasEdge(1, 2) {
			t.Error("expected false for non-existent edge")
		}
	})
}

func TestGraph_Neighbors(t *testing.T) {
	t.Run("Neighbors returns correct list for undirected graph", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(1, 4)

		neighbors := g.Neighbors(1)
		sort.Ints(neighbors)
		expected := []int{2, 3, 4}
		if !reflect.DeepEqual(neighbors, expected) {
			t.Errorf("expected neighbors %v, got %v", expected, neighbors)
		}
	})

	t.Run("Neighbors returns correct list for directed graph", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(4, 1) // 4->1, so 4 should NOT be a neighbor of 1

		neighbors := g.Neighbors(1)
		sort.Ints(neighbors)
		expected := []int{2, 3}
		if !reflect.DeepEqual(neighbors, expected) {
			t.Errorf("expected neighbors %v, got %v", expected, neighbors)
		}
	})

	t.Run("Neighbors returns nil for non-existent vertex", func(t *testing.T) {
		g := NewGraph(false)
		neighbors := g.Neighbors(99)
		if neighbors != nil {
			t.Errorf("expected nil for non-existent vertex, got %v", neighbors)
		}
	})
}

// Build a deterministic graph for traversal tests:
//
//	1 --- 2 --- 5
//	|     |
//	3 --- 4
func buildTestGraph() *Graph {
	g := NewGraph(false)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	return g
}

func TestGraph_BFS(t *testing.T) {
	t.Run("BFS visit order on known graph", func(t *testing.T) {
		g := buildTestGraph()
		order := g.BFS(1)

		// With sorted neighbors: 1 -> [2,3] -> [4,5] (from 2), [4 already visited] (from 3)
		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(order, expected) {
			t.Errorf("expected BFS order %v, got %v", expected, order)
		}
	})

	t.Run("BFS from non-existent vertex returns nil", func(t *testing.T) {
		g := buildTestGraph()
		order := g.BFS(99)
		if order != nil {
			t.Errorf("expected nil, got %v", order)
		}
	})

	t.Run("BFS on single vertex", func(t *testing.T) {
		g := NewGraph(false)
		g.AddVertex(1)
		order := g.BFS(1)
		expected := []int{1}
		if !reflect.DeepEqual(order, expected) {
			t.Errorf("expected %v, got %v", expected, order)
		}
	})
}

func TestGraph_DFS(t *testing.T) {
	t.Run("DFS visit order on known graph", func(t *testing.T) {
		g := buildTestGraph()
		order := g.DFS(1)

		// With sorted neighbors: 1 -> 2 -> 4 -> 3 -> ... -> 5
		expected := []int{1, 2, 4, 3, 5}
		if !reflect.DeepEqual(order, expected) {
			t.Errorf("expected DFS order %v, got %v", expected, order)
		}
	})

	t.Run("DFS from non-existent vertex returns nil", func(t *testing.T) {
		g := buildTestGraph()
		order := g.DFS(99)
		if order != nil {
			t.Errorf("expected nil, got %v", order)
		}
	})

	t.Run("DFS on single vertex", func(t *testing.T) {
		g := NewGraph(false)
		g.AddVertex(1)
		order := g.DFS(1)
		expected := []int{1}
		if !reflect.DeepEqual(order, expected) {
			t.Errorf("expected %v, got %v", expected, order)
		}
	})
}

func TestGraph_ShortestPath(t *testing.T) {
	t.Run("ShortestPath between connected nodes", func(t *testing.T) {
		g := buildTestGraph()
		path, dist := g.ShortestPath(1, 5)

		expectedPath := []int{1, 2, 5}
		expectedDist := 2
		if !reflect.DeepEqual(path, expectedPath) {
			t.Errorf("expected path %v, got %v", expectedPath, path)
		}
		if dist != expectedDist {
			t.Errorf("expected distance %d, got %d", expectedDist, dist)
		}
	})

	t.Run("ShortestPath to self returns distance 0", func(t *testing.T) {
		g := buildTestGraph()
		path, dist := g.ShortestPath(1, 1)

		expectedPath := []int{1}
		if !reflect.DeepEqual(path, expectedPath) {
			t.Errorf("expected path %v, got %v", expectedPath, path)
		}
		if dist != 0 {
			t.Errorf("expected distance 0, got %d", dist)
		}
	})

	t.Run("ShortestPath between disconnected nodes returns nil and -1", func(t *testing.T) {
		g := NewGraph(false)
		g.AddEdge(1, 2)
		g.AddVertex(3) // isolated vertex

		path, dist := g.ShortestPath(1, 3)
		if path != nil {
			t.Errorf("expected nil path, got %v", path)
		}
		if dist != -1 {
			t.Errorf("expected distance -1, got %d", dist)
		}
	})

	t.Run("ShortestPath with non-existent start vertex", func(t *testing.T) {
		g := buildTestGraph()
		path, dist := g.ShortestPath(99, 1)
		if path != nil {
			t.Errorf("expected nil path, got %v", path)
		}
		if dist != -1 {
			t.Errorf("expected distance -1, got %d", dist)
		}
	})

	t.Run("ShortestPath with non-existent end vertex", func(t *testing.T) {
		g := buildTestGraph()
		path, dist := g.ShortestPath(1, 99)
		if path != nil {
			t.Errorf("expected nil path, got %v", path)
		}
		if dist != -1 {
			t.Errorf("expected distance -1, got %d", dist)
		}
	})

	t.Run("ShortestPath in directed graph", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(1, 3) // direct edge 1->3

		path, dist := g.ShortestPath(1, 3)
		expectedPath := []int{1, 3}
		if !reflect.DeepEqual(path, expectedPath) {
			t.Errorf("expected path %v, got %v", expectedPath, path)
		}
		if dist != 1 {
			t.Errorf("expected distance 1, got %d", dist)
		}
	})

	t.Run("ShortestPath in directed graph with no reverse path", func(t *testing.T) {
		g := NewGraph(true)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)

		path, dist := g.ShortestPath(3, 1)
		if path != nil {
			t.Errorf("expected nil path, got %v", path)
		}
		if dist != -1 {
			t.Errorf("expected distance -1, got %d", dist)
		}
	})
}
