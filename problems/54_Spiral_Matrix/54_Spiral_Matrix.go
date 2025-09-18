package problems

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	n := len(matrix)
	m := len(matrix[0])
	top, bottom := 0, n-1
	left, right := 0, m-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			// left top -> right top
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			// right top -> right bottom
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				// right bottom -> left bottom
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				// left bottom -> left top
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}
