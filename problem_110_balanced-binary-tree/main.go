package problem_110_balanced_binary_tree

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

	return depth(root) != -1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := depth(node.Left)
	if left == -1 {
		return -1
	}

	right := depth(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left)-float64(right)) < 2 {
		return int(math.Max(float64(left), float64(right)) + 1)
	} else {
		return -1
	}
}
