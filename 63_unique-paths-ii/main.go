package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, cols)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[len(dp)-1]
}
