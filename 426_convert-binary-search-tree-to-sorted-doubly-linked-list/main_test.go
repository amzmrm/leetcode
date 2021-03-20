package leetcode

import "testing"

func TestTreeToDoublyList(t *testing.T) {
	root := &Node{Val: 4}
	node1 := &Node{Val: 2}
	node2 := &Node{Val: 5}
	node3 := &Node{Val: 1}
	node4 := &Node{Val: 3}
	root.Left, root.Right = node1, node2
	node1.Left, node1.Right = node3, node4

	got := treeToDoublyList(root)
	_ = got
}
