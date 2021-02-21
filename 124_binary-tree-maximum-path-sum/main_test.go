package leetcode

import "testing"

func TestMaxPathSum(t *testing.T) {
	root := TreeNode{Val: -10}
	node1 := TreeNode{Val: 9}
	node2 := TreeNode{Val: 20}
	node3 := TreeNode{Val: 15}
	node4 := TreeNode{Val: 7}

	root.Left = &node1
	root.Right = &node2
	node2.Left = &node3
	node2.Right = &node4

	got := MaxPathSum(&root)
	expect := 42
	if got != expect {
		t.Fatalf("expect: %d, got: %d", expect, got)
	}
}
