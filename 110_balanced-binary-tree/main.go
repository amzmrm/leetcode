package leetcode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := height(node.Left)
	if left == -1 {
		return -1
	}

	right := height(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left)-float64(right)) < 2 {
		return int(math.Max(float64(left), float64(right)) + 1)
	} else {
		return -1
	}
}
