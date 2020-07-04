package main

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	t.Run("both on lift sub-tree of root", func(t *testing.T) {
		root := TreeNode{Val: 6}
		node2 := TreeNode{Val: 2}
		node3 := TreeNode{Val: 8}
		node4 := TreeNode{Val: 0}
		node5 := TreeNode{Val: 4}
		node6 := TreeNode{Val: 3}
		node7 := TreeNode{Val: 5}
		node8 := TreeNode{Val: 7}
		node9 := TreeNode{Val: 9}
		root.Left = &node2
		root.Right = &node3
		node2.Left = &node4
		node2.Right = &node5
		node5.Left = &node6
		node5.Right = &node7
		node3.Left = &node8
		node3.Right = &node9

		got := LowestCommonAncestor(&root, &node2, &node5)
		want := 2
		if got.Val != want {
			t.Fatalf("got:%v, expect:%v", got, want)
		}
	})

	t.Run("both on right sub-tree of root", func(t *testing.T) {
		root := TreeNode{Val: 6}
		node2 := TreeNode{Val: 2}
		node3 := TreeNode{Val: 8}
		node4 := TreeNode{Val: 0}
		node5 := TreeNode{Val: 4}
		node6 := TreeNode{Val: 3}
		node7 := TreeNode{Val: 5}
		node8 := TreeNode{Val: 7}
		node9 := TreeNode{Val: 9}
		root.Left = &node2
		root.Right = &node3
		node2.Left = &node4
		node2.Right = &node5
		node5.Left = &node6
		node5.Right = &node7
		node3.Left = &node8
		node3.Right = &node9

		got := LowestCommonAncestor(&root, &node8, &node9)
		want := 8
		if got.Val != want {
			t.Fatalf("got:%v, expect:%v", got, want)
		}
	})

	t.Run("one node on the lift sub-tree, another one on the right sub-tree", func(t *testing.T) {
		root := TreeNode{Val: 6}
		node2 := TreeNode{Val: 2}
		node3 := TreeNode{Val: 8}
		node4 := TreeNode{Val: 0}
		node5 := TreeNode{Val: 4}
		node6 := TreeNode{Val: 3}
		node7 := TreeNode{Val: 5}
		node8 := TreeNode{Val: 7}
		node9 := TreeNode{Val: 9}
		root.Left = &node2
		root.Right = &node3
		node2.Left = &node4
		node2.Right = &node5
		node5.Left = &node6
		node5.Right = &node7
		node3.Left = &node8
		node3.Right = &node9

		got := LowestCommonAncestor(&root, &node4, &node9)
		want := 6
		if got.Val != want {
			t.Fatalf("got:%v, expect:%v", got, want)
		}
	})
}
