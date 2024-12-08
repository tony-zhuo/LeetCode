package problems

func exist(board [][]byte, word string) bool {
	for m := 0; m < len(board); m++ {
		for n := 0; n < len(board[0]); n++ {
			if dfs(board, m, n, 0, word) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, m, n, index int, word string) bool {
	if len(word) == index {
		return true
	}
	if m < 0 || m >= len(board) || n < 0 || n >= len(board[m]) || board[m][n] != word[index] {
		return false
	}

	tmp := board[m][n]
	board[m][n] = '*'
	found := dfs(board, m+1, n, index+1, word) ||
		dfs(board, m-1, n, index+1, word) ||
		dfs(board, m, n+1, index+1, word) ||
		dfs(board, m, n-1, index+1, word)

	board[m][n] = tmp

	return found
}
