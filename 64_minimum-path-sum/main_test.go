package leetcode

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	res := minPathSumWithDP2(grid)
	t.Log(res)
}
