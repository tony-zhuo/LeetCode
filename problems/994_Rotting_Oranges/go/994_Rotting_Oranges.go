package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap

// use bfs
// 先把壞橘子放到 queue 裡面
// 再從 queue 把壞橘子拿出來，如果旁邊有好橘子，把他變壞橘子之後放到 queue 裡面
// 跑完比對壞橘子數量、總橘子數量
func orangesRotting(grid [][]int) int {
	queue := [][]int{}
	freshOranges := 0

	// 1. 初始化 Queue 並統計新鮮橘子的數量
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	if freshOranges == 0 {
		return 0
	}

	minutes := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		size := len(queue)
		rottenThisMinute := false

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			x, y := curr[0], curr[1]

			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]

				if nx >= 0 &&
					nx < len(grid) &&
					ny >= 0 &&
					ny < len(grid[0]) &&
					grid[nx][ny] == 1 {

					grid[nx][ny] = 2
					queue = append(queue, []int{nx, ny})
					freshOranges--
					rottenThisMinute = true
				}
			}
		}

		if rottenThisMinute {
			minutes++
		}

	}

	// 3. 檢查是否還有無法被感染的新鮮橘子
	if freshOranges > 0 {
		return -1
	}

	return minutes
}
