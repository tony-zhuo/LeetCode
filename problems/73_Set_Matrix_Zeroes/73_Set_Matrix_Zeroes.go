package problems

//func setZeroes(matrix [][]int) {
//	row := len(matrix)
//	column := len(matrix[0])
//
//	rows_zero := make([]bool, row)
//	columns_zero := make([]bool, column)
//
//	for i := 0; i <= row-1; i++ {
//		for j := 0; j <= column-1; j++ {
//			if matrix[i][j] == 0 {
//				rows_zero[i] = true
//				columns_zero[j] = true
//			}
//		}
//	}
//
//	for index, item_isZero := range rows_zero {
//		if item_isZero {
//			for i := 0; i < len(matrix[index]); i++ {
//				matrix[index][i] = 0
//			}
//		}
//	}
//
//	for index, item_isZero := range columns_zero {
//		if item_isZero {
//			for i := 0; i < len(matrix); i++ {
//				matrix[i][index] = 0
//			}
//		}
//	}
//}

func setZeroes(matrix [][]int) {
	row := len(matrix) - 1
	column := len(matrix[0]) - 1

	is_first_rows_zero := false
	is_first_columns_zero := false

	for i := 0; i <= row; i++ {
		if matrix[i][0] == 0 {
			is_first_columns_zero = true
			break
		}
	}

	for i := 0; i <= column; i++ {
		if matrix[0][i] == 0 {
			is_first_rows_zero = true
			break
		}
	}

	// use first row and first column as markers
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// set zero based on markers
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if is_first_rows_zero {
		for j := 0; j <= column; j++ {
			matrix[0][j] = 0
		}
	}

	if is_first_columns_zero {
		for i := 0; i <= row; i++ {
			matrix[i][0] = 0
		}
	}
}
