package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prev, target1, target2 *TreeNode

	var rec func(node *TreeNode)
	rec = func(node *TreeNode) {
		if node == nil {
			return
		}

		rec(node.Left)

		if prev != nil && prev.Val > node.Val {
			if target1 == nil {
				target1 = prev
			}
			target2 = node
		}
		prev = node

		rec(node.Right)
	}

	rec(root)

	if target1 != nil && target2 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}
}
