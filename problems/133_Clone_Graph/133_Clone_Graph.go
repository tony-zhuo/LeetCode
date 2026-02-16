package problems

import datastructures "leet-code/data_structures/graph"

func cloneGraph(node *datastructures.Node) *datastructures.Node {
	if node == nil {
		return nil
	}
	visited := make(map[*datastructures.Node]*datastructures.Node)

	return dfs(node, visited)
}

func dfs(node *datastructures.Node, visited map[*datastructures.Node]*datastructures.Node) *datastructures.Node {
	if n, ok := visited[node]; ok {
		return n
	}
	tmp := &datastructures.Node{Val: node.Val}
	visited[node] = tmp
	for _, n := range node.Neighbors {
		tmp.Neighbors = append(tmp.Neighbors, dfs(n, visited))
	}
	return tmp
}
