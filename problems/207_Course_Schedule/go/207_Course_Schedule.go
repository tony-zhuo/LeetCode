package problems

func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, prerequisite := range prerequisites {
		x := prerequisite[0] // 想修的課
		y := prerequisite[1] // 修想修的課，必須先修的課
		graph[x] = append(graph[x], y)
	}

	visited := make(map[int]int, numCourses)
	// 0 -> init
	// 1 -> has cycle
	// -1 -> alreafy visited but no cycle
	for i := 0; i < numCourses; i++ {
		if dfs(graph, visited, i) {
			return false
		}
	}

	return true
}

func dfs(graph [][]int, visited map[int]int, numCourse int) bool {
	if visited[numCourse] == 1 {
		return true
	}
	if visited[numCourse] == -1 {
		return false
	}

	visited[numCourse] = 1

	for _, child := range graph[numCourse] {
		if dfs(graph, visited, child) {
			return true
		}
	}

	visited[numCourse] = -1

	return false
}
