package datastructures

import "sort"

// Graph using adjacency list representation
type Graph struct {
	adjacency map[int][]int
	directed  bool
}

// NewGraph creates a new graph. If directed is true, edges are one-directional.
func NewGraph(directed bool) *Graph {
	return &Graph{
		adjacency: make(map[int][]int),
		directed:  directed,
	}
}

// AddVertex adds a vertex with an empty adjacency list if it does not already exist.
func (g *Graph) AddVertex(v int) {
	if _, exists := g.adjacency[v]; !exists {
		g.adjacency[v] = []int{}
	}
}

// AddEdge adds an edge from 'from' to 'to'. For undirected graphs, the reverse
// edge is also added. Vertices are auto-created if they do not exist.
func (g *Graph) AddEdge(from, to int) {
	g.AddVertex(from)
	g.AddVertex(to)
	g.adjacency[from] = append(g.adjacency[from], to)
	if !g.directed {
		g.adjacency[to] = append(g.adjacency[to], from)
	}
}

// RemoveEdge removes the edge from 'from' to 'to'. For undirected graphs, the
// reverse edge is also removed.
func (g *Graph) RemoveEdge(from, to int) {
	g.adjacency[from] = removeFromSlice(g.adjacency[from], to)
	if !g.directed {
		g.adjacency[to] = removeFromSlice(g.adjacency[to], from)
	}
}

// removeFromSlice removes the first occurrence of val from the slice.
func removeFromSlice(slice []int, val int) []int {
	for i, v := range slice {
		if v == val {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// HasEdge returns true if an edge exists from 'from' to 'to'.
func (g *Graph) HasEdge(from, to int) bool {
	neighbors, exists := g.adjacency[from]
	if !exists {
		return false
	}
	for _, v := range neighbors {
		if v == to {
			return true
		}
	}
	return false
}

// Neighbors returns the list of adjacent vertices for a given vertex.
func (g *Graph) Neighbors(v int) []int {
	return g.adjacency[v]
}

// Vertices returns a sorted list of all vertices in the graph.
func (g *Graph) Vertices() []int {
	vertices := make([]int, 0, len(g.adjacency))
	for v := range g.adjacency {
		vertices = append(vertices, v)
	}
	sort.Ints(vertices)
	return vertices
}

// BFS performs a breadth-first search starting from the given vertex and returns
// the visit order.
func (g *Graph) BFS(start int) []int {
	if _, exists := g.adjacency[start]; !exists {
		return nil
	}

	visited := make(map[int]bool)
	order := []int{}
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		order = append(order, current)

		neighbors := make([]int, len(g.adjacency[current]))
		copy(neighbors, g.adjacency[current])
		sort.Ints(neighbors)

		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return order
}

// DFS performs a depth-first search starting from the given vertex and returns
// the visit order.
func (g *Graph) DFS(start int) []int {
	if _, exists := g.adjacency[start]; !exists {
		return nil
	}

	visited := make(map[int]bool)
	order := []int{}
	g.dfsHelper(start, visited, &order)
	return order
}

func (g *Graph) dfsHelper(v int, visited map[int]bool, order *[]int) {
	visited[v] = true
	*order = append(*order, v)

	neighbors := make([]int, len(g.adjacency[v]))
	copy(neighbors, g.adjacency[v])
	sort.Ints(neighbors)

	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited, order)
		}
	}
}

// ShortestPath finds the shortest path between start and end using BFS (for
// unweighted graphs). It returns the path as a slice of vertices and the
// distance. If no path exists, it returns (nil, -1).
func (g *Graph) ShortestPath(start, end int) ([]int, int) {
	if _, exists := g.adjacency[start]; !exists {
		return nil, -1
	}
	if _, exists := g.adjacency[end]; !exists {
		return nil, -1
	}
	if start == end {
		return []int{start}, 0
	}

	visited := make(map[int]bool)
	parent := make(map[int]int)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.adjacency[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = current
				if neighbor == end {
					return reconstructPath(parent, start, end)
				}
				queue = append(queue, neighbor)
			}
		}
	}

	return nil, -1
}

// reconstructPath builds the path from start to end using the parent map.
func reconstructPath(parent map[int]int, start, end int) ([]int, int) {
	path := []int{}
	current := end
	for current != start {
		path = append(path, current)
		current = parent[current]
	}
	path = append(path, start)

	// Reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, len(path) - 1
}
