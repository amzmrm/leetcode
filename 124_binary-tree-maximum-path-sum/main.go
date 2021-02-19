package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(maxGain(node.Left), 0)
		right := max(maxGain(node.Right), 0)

		priceNewPath := node.Val + left + right

		maxSum = max(maxSum, priceNewPath)

		return node.Val + max(left, right)
	}

	maxGain(root)

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
