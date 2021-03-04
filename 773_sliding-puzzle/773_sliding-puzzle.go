package leetcode

import "strconv"

func slidingPuzzle(board [][]int) int {
	var start string
	for _, nums := range board {
		for _, num := range nums {
			start += strconv.Itoa(num)
		}
	}
	target := "123450"

	neighbours := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}

	var ascii0 uint8 = 48
	var queue []string
	queue = append(queue, start)
	visited := make(map[string]bool)
	visited[start] = true
	var step int
	var curr string
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr, queue = queue[0], queue[1:]
			if curr == target {
				return step
			}

			var idx int
			for j := 0; j < len(curr); j++ {
				if curr[j] != ascii0 {
					continue
				}
				idx = j
			}

			for _, adj := range neighbours[idx] {
				newBoard := []rune(curr)
				newBoard[adj], newBoard[idx] = newBoard[idx], newBoard[adj]
				newBoardStr := string(newBoard)
				if _, ok := visited[newBoardStr]; !ok {
					queue = append(queue, newBoardStr)
					visited[newBoardStr] = true
				}
			}
		}
		step++
	}
	return -1
}
