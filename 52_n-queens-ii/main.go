package main

var asciiDot uint8 = 46
var asciiQ uint8 = 81

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var res [][]string
	board := make([]string, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i] += "."
		}
	}

	backtrack(&board, &res, 0)

	return len(res)
}

func backtrack(board *[]string, res *[][]string, row int) {
	if len(*board) == row {
		newBoard := make([]string, len(*board))
		copy(newBoard, *board)
		*res = append(*res, newBoard)
		return
	}

	n := len((*board)[row])
	for col := 0; col < n; col++ {
		if !isValid(*board, row, col) {
			continue
		}

		(*board)[row] = (*board)[row][:col] + "Q" + (*board)[row][col+1:]
		backtrack(board, res, row+1)
		(*board)[row] = (*board)[row][:col] + "." + (*board)[row][col+1:]
	}
}

func isValid(board []string, row int, col int) bool {
	n := len(board)

	// 检查col列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i][col] == asciiQ {
			return false
		}
	}

	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; {
		if board[i][j] == asciiQ {
			return false
		}
		i--
		j++
	}

	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == asciiQ {
			return false
		}
		i--
		j--
	}

	return true
}
