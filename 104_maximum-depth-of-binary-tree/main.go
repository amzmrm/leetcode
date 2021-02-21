package leetcode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归公式：maxDepth(root) = 1 + max(depth(root.Left), depth(root.Right))
// 递归结束条件：root == nil
// 时间复杂度：O(n)
// 空间复杂度：最好：O(logn), 最坏：O(n)
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func MaxDepthIterative(root *TreeNode) int {
	type pair struct {
		node  *TreeNode
		depth int
	}

	stack := make([]pair, 0)
	stack = append(stack, pair{
		node:  root,
		depth: 1,
	})

	currentDepth, maxDepth := 0, 0
	var currNode pair
	for len(stack) > 0 {
		currNode, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		if currNode.node == nil {
			continue
		}

		currentDepth = currNode.depth
		if maxDepth < currentDepth {
			maxDepth = currentDepth
		}

		stack = append(stack, pair{
			node:  currNode.node.Left,
			depth: currentDepth + 1,
		})
		stack = append(stack, pair{
			node:  currNode.node.Right,
			depth: currentDepth + 1,
		})
	}

	return maxDepth
}
