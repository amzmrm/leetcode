package leetcode

func solveSudoku(board [][]byte) {
	backtrack(&board, 0, 0)
}

func backtrack(board *[][]byte, row int, col int) bool {
	if col == len(*board) {
		return backtrack(board, row+1, 0)
	}
	if row == len(*board) {
		return true
	}
	if (*board)[row][col] != byte('.') {
		return backtrack(board, row, col+1)
	}
	for ch := byte('1'); ch <= byte('9'); ch++ {
		if !isValid(board, row, col, ch) {
			continue
		}
		(*board)[row][col] = ch
		if backtrack(board, row, col+1) {
			return true
		}
		(*board)[row][col] = byte('.')
	}
	return false
}

func isValid(board *[][]byte, row int, col int, n byte) bool {
	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if (*board)[row][i] == n {
			return false
		}
		// 判断列是否存在重复
		if (*board)[i][col] == n {
			return false
		}
		// 判断 3 x 3 方框是否存在重复
		if (*board)[(row/3)*3+i/3][(col/3)*3+i%3] == n {
			return false
		}
	}
	return true
}
