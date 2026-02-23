package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	for i := range mat {
		res[i] = make([]int, len(mat[i]))
	}

	queue := [][2]int{}
	for x := range mat {
		for y := range mat[x] {
			if mat[x][y] == 0 {
				queue = append(queue, [2]int{x, y})
				res[x][y] = 0
			} else {
				res[x][y] = -1
			}
		}
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			xn, yn := curr[0]+dir[0], curr[1]+dir[1]
			if xn < 0 || xn >= len(res) || yn < 0 || yn >= len(res[xn]) || res[xn][yn] != -1 {
				continue
			}

			res[xn][yn] = res[curr[0]][curr[1]] + 1
			queue = append(queue, [2]int{xn, yn})
		}
	}

	return res
}
