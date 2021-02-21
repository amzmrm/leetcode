package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left, right := root.Left, root.Right
	root.Left, root.Right = nil, left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
