package problems

import "leet-code/structure"

func cloneGraph(node *structure.Node) *structure.Node {
	if node == nil {
		return nil
	}
	visited := make(map[*structure.Node]*structure.Node)

	return dfs(node, visited)
}

func dfs(node *structure.Node, visited map[*structure.Node]*structure.Node) *structure.Node {
	if n, ok := visited[node]; ok {
		return n
	}
	tmp := &structure.Node{Val: node.Val}
	visited[node] = tmp
	for _, n := range node.Neighbors {
		tmp.Neighbors = append(tmp.Neighbors, dfs(n, visited))
	}
	return tmp
}
