package leetcode

func ClimbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	results := []int{0, 1, 2}
	if n <= 2 {
		return results[n]
	}

	stairOne, stairTwo, stairN := 2, 1, 0
	for i := 3; i <= n; i++ {
		stairN = stairOne + stairTwo
		stairTwo, stairOne = stairOne, stairN
	}

	return stairN
}
