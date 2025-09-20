package problems

func setZeroes(matrix [][]int) {
	row := len(matrix)
	column := len(matrix[0])

	rows_zero := make([]bool, row)
	columns_zero := make([]bool, column)

	for i := 0; i <= row-1; i++ {
		for j := 0; j <= column-1; j++ {
			if matrix[i][j] == 0 {
				rows_zero[i] = true
				columns_zero[j] = true
			}
		}
	}

	for index, item_isZero := range rows_zero {
		if item_isZero {
			for i := 0; i < len(matrix[index]); i++ {
				matrix[index][i] = 0
			}
		}
	}

	for index, item_isZero := range columns_zero {
		if item_isZero {
			for i := 0; i < len(matrix); i++ {
				matrix[i][index] = 0
			}
		}
	}
}
