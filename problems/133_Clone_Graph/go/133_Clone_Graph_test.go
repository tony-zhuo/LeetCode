package problems

import (
	datastructures "leet-code/data_structures/graph/go"
	"testing"
)

// buildGraph builds an undirected graph from an adjacency list.
// adjList[i] contains the 1-indexed neighbor values for node i+1.
func buildGraph(adjList [][]int) *datastructures.Node {
	if len(adjList) == 0 {
		return nil
	}
	nodes := make(map[int]*datastructures.Node)
	for i := 1; i <= len(adjList); i++ {
		nodes[i] = &datastructures.Node{Val: i}
	}
	for i, neighbors := range adjList {
		for _, v := range neighbors {
			nodes[i+1].Neighbors = append(nodes[i+1].Neighbors, nodes[v])
		}
	}
	return nodes[1]
}

// toAdjList converts a graph (starting from node) into a sorted adjacency list.
func toAdjList(node *datastructures.Node) [][]int {
	if node == nil {
		return nil
	}
	visited := make(map[int]*datastructures.Node)
	var walk func(n *datastructures.Node)
	walk = func(n *datastructures.Node) {
		if _, ok := visited[n.Val]; ok {
			return
		}
		visited[n.Val] = n
		for _, nb := range n.Neighbors {
			walk(nb)
		}
	}
	walk(node)

	result := make([][]int, len(visited))
	for i := 1; i <= len(visited); i++ {
		n := visited[i]
		neighbors := make([]int, len(n.Neighbors))
		for j, nb := range n.Neighbors {
			neighbors[j] = nb.Val
		}
		result[i-1] = neighbors
	}
	return result
}

var cases = []struct {
	name    string
	adjList [][]int
}{
	{
		name:    "example 1: 4-node graph",
		adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
	},
	{
		name:    "example 2: single node",
		adjList: [][]int{{}},
	},
	{
		name:    "example 3: empty graph",
		adjList: [][]int{},
	},
	{
		name:    "two connected nodes",
		adjList: [][]int{{2}, {1}},
	},
}

func Test_cloneGraph(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			cloned := cloneGraph(original)

			if original == nil {
				if cloned != nil {
					t.Errorf("expected nil, got non-nil")
				}
				return
			}

			// Verify structure matches
			origAdj := toAdjList(original)
			clonedAdj := toAdjList(cloned)
			if len(origAdj) != len(clonedAdj) {
				t.Errorf("node count mismatch: original %d, cloned %d", len(origAdj), len(clonedAdj))
				return
			}
			for i := range origAdj {
				if len(origAdj[i]) != len(clonedAdj[i]) {
					t.Errorf("node %d neighbor count mismatch", i+1)
					return
				}
				for j := range origAdj[i] {
					if origAdj[i][j] != clonedAdj[i][j] {
						t.Errorf("node %d neighbor %d: got %d, want %d", i+1, j, clonedAdj[i][j], origAdj[i][j])
					}
				}
			}

			// Verify it's a deep copy (different pointers)
			if cloned == original {
				t.Errorf("cloned node is the same pointer as original")
			}
		})
	}
}