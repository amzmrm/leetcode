package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	trees := make(map[string]int)

	var res []*TreeNode

	var fn func(root *TreeNode) string
	fn = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}

		left := fn(root.Left)
		right := fn(root.Right)

		subTree := fmt.Sprintf("%v,%v,%v", left, right, root.Val)

		if val, ok := trees[subTree]; ok {
			if val == 1 {
				res = append(res, root)
			}
			trees[subTree] = val + 1
		} else {
			trees[subTree] = 1
		}

		return subTree
	}

	fn(root)

	return res
}
