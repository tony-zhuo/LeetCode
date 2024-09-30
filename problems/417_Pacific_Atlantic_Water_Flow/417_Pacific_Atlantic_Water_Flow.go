package problems

import "math"

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}
	row, col, res := len(heights), len(heights[0]), make([][]int, 0)
	pacific := make([][]bool, row)
	atlantic := make([][]bool, row)
	for i := 0; i < row; i++ {
		pacific[i] = make([]bool, col)
		atlantic[i] = make([]bool, col)
	}

	for r := 0; r < row; r++ {
		dfs(r, 0, math.MinInt, heights, &pacific)
		dfs(r, col-1, math.MinInt, heights, &atlantic)
	}
	for c := 0; c < col; c++ {
		dfs(0, c, math.MinInt, heights, &pacific)
		dfs(row-1, c, math.MinInt, heights, &atlantic)
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

func dfs(row, col, height int, heights [][]int, ocean *[][]bool) {
	// 邊界檢查
	if row < 0 || row >= len(heights) || col < 0 || col >= len(heights[0]) {
		return
	}
	// 是否被訪問過
	if (*ocean)[row][col] {
		return
	}
	// 高度檢查（水只能從低點流向高點，這邊是從低處開始）
	if heights[row][col] < height {
		return
	}
	(*ocean)[row][col] = true
	dfs(row+1, col, heights[row][col], heights, ocean)
	dfs(row-1, col, heights[row][col], heights, ocean)
	dfs(row, col+1, heights[row][col], heights, ocean)
	dfs(row, col-1, heights[row][col], heights, ocean)
}
