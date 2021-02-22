package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var rank int
	var res *TreeNode
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)

		rank++
		if rank == k {
			res = root
			return
		}

		traverse(root.Right)
	}

	traverse(root)

	return res.Val
}
