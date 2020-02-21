package problem_4_find_num_in_2d_array

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, columns := len(matrix), len(matrix[0])
	row, column := 0, columns-1
	num := 0
	for row < rows && column >= 0 {
		num = matrix[row][column]
		if num == target {
			return true
		} else if num > target {
			column--
		} else {
			row++
		}
	}
	return false
}
