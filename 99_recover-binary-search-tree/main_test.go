package leetcode

import "testing"

func TestBuildTree(t *testing.T) {
	root := TreeNode{Val: 2}
	node1 := TreeNode{Val: 3}
	node2 := TreeNode{Val: 1}

	root.Left = &node1
	root.Right = &node2

	recoverTree(&root)
	t.Log(root.Val)
}
