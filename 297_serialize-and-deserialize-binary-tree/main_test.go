package leetcode

import "testing"

func TestSel(t *testing.T) {
	root := TreeNode{Val: 1}
	node1 := TreeNode{Val: 2}
	node2 := TreeNode{Val: 3}
	node3 := TreeNode{Val: 4}
	node4 := TreeNode{Val: 5}
	root.Left = &node1
	root.Right = &node2
	node2.Left = &node3
	node2.Right = &node4

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(&root)
	got := deser.deserialize(data)
	if got.Val != root.Val {
		t.Fatalf("expect: %d, got: %d", root.Val, got.Val)
	}
}
