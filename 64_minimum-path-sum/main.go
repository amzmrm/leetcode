package leetcode

import "math"

func minPathSum(grid [][]int) int {
	res := math.MaxInt32
	backtrack(grid, 0, 0, grid[0][0], &res)
	return res
}

func minPathSumWithMemo(grid [][]int) int {
	memo := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		memo[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			memo[i][j] = math.MaxInt32
		}
	}
	return backtrackWithMemo(grid, 0, 0, &memo)
}

func backtrack(grid [][]int, row, col int, sum int, res *int) {
	if row == len(grid)-1 && col == len(grid[0])-1 {
		if sum < *res {
			*res = sum
		}
		return
	}
	dirs := [][]int{{1, 0}, {0, 1}}
	for _, dir := range dirs {
		nextRow := row + dir[0]
		nextCol := col + dir[1]
		if nextRow >= len(grid) || nextCol >= len(grid[0]) {
			continue
		}
		sum += grid[nextRow][nextCol]
		backtrack(grid, nextRow, nextCol, sum, res)
		sum -= grid[nextRow][nextCol]
	}
}

func backtrackWithMemo(grid [][]int, row, col int, memo *[][]int) int {
	if row == len(grid)-1 && col == len(grid[0])-1 {
		return grid[row][col]
	}
	if (*memo)[row][col] != math.MaxInt32 {
		return (*memo)[row][col]
	}

	min := math.MaxInt32
	dirs := [][]int{{1, 0}, {0, 1}}
	for _, dir := range dirs {
		nextRow := row + dir[0]
		nextCol := col + dir[1]
		if nextRow >= len(grid) || nextCol >= len(grid[0]) {
			continue
		}
		childPathSum := backtrackWithMemo(grid, nextRow, nextCol, memo)
		if childPathSum < min {
			min = childPathSum
		}
	}
	(*memo)[row][col] = min + grid[row][col]
	return (*memo)[row][col]
}

func minPathSumWithDP(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 准备好states状态数组
	rows, cols := len(grid), len(grid[0])
	states := make([][]int, rows)
	for i := 0; i < rows; i++ {
		states[i] = make([]int, cols)
	}

	// 初始化states的第一列数据
	sum := 0
	for i := 0; i < rows; i++ {
		sum += grid[i][0]
		states[i][0] = sum
	}
	// 初始化states的第一行数据
	sum = 0
	for j := 0; j < cols; j++ {
		sum += grid[0][j]
		states[0][j] = sum
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			states[i][j] = grid[i][j] + min(states[i-1][j], states[i][j-1])
		}
	}

	return states[rows-1][cols-1]
}

func minPathSumWithDP2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	states := make([]int, cols)
	states[0] = grid[0][0] // 初始化状态数组第一个元素
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j != 0 {
				states[j] = grid[i][j] + states[j-1]
			} else if i != 0 && j == 0 {
				states[j] = grid[i][j] + states[j]
			} else if i != 0 && j != 0 {
				states[j] = grid[i][j] + min(states[j], states[j-1])
			}
		}
	}
	return states[cols-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
