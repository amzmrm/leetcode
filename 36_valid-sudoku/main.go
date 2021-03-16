package leetcode

import "strconv"

func isValidSudoku(board [][]byte) bool {
	rows := make([][]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]int, 9)
	}
	cols := make([][]int, 9)
	for i := 0; i < 9; i++ {
		cols[i] = make([]int, 9)
	}
	box := make([][]int, 9)
	for i := 0; i < 9; i++ {
		box[i] = make([]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == byte('.') {
				continue
			}
			val, _ := strconv.Atoi(string(board[i][j]))
			val = val - 1
			boxIdx := (i/3)*3 + j/3
			if rows[i][val] == 0 && cols[j][val] == 0 && box[boxIdx][val] == 0 {
				rows[i][val] = 1
				cols[j][val] = 1
				box[boxIdx][val] = 1
			} else {
				return false
			}
		}
	}
	return true
}
