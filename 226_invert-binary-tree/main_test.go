package leetcode

import "testing"

func TestInvertTree(t *testing.T) {
	root := TreeNode{Val: 4}
	node1 := TreeNode{Val: 2}
	node2 := TreeNode{Val: 7}
	node3 := TreeNode{Val: 1}
	node4 := TreeNode{Val: 3}
	node5 := TreeNode{Val: 6}
	node6 := TreeNode{Val: 9}
	root.Left = &node1
	root.Right = &node2
	node1.Left = &node3
	node1.Right = &node4
	node2.Left = &node5
	node2.Right = &node6

	got := invertTree(&root)
	_ = got
}
