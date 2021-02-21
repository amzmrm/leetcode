package leetcode

import "testing"

func TestLowestCommonAncestor1(t *testing.T) {

}

func TestLowestCommonAncestor2(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		root := TreeNode{Val: 3}
		node1 := TreeNode{Val: 5}
		node2 := TreeNode{Val: 6}
		node3 := TreeNode{Val: 2}
		node4 := TreeNode{Val: 7}
		node5 := TreeNode{Val: 4}
		node6 := TreeNode{Val: 1}
		node7 := TreeNode{Val: 0}
		node8 := TreeNode{Val: 8}

		root.Left = &node1
		node1.Left = &node2
		node1.Right = &node3
		node3.Left = &node4
		node3.Right = &node5

		root.Right = &node6
		node6.Left = &node7
		node6.Right = &node8

		p := TreeNode{Val: 5}
		q := TreeNode{Val: 1}
		expect := TreeNode{Val: 3}
		got := LowestCommonAncestor2(&root, &p, &q)
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}

		q = TreeNode{Val: 5}
		q = TreeNode{Val: 4}
		expect = TreeNode{Val: 5}
		got = LowestCommonAncestor2(&root, &p, &q)
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}
	})

	t.Run("recursively", func(t *testing.T) {
		root := TreeNode{Val: 3}
		node1 := TreeNode{Val: 5}
		node2 := TreeNode{Val: 6}
		node3 := TreeNode{Val: 2}
		node4 := TreeNode{Val: 7}
		node5 := TreeNode{Val: 4}
		node6 := TreeNode{Val: 1}
		node7 := TreeNode{Val: 0}
		node8 := TreeNode{Val: 8}

		root.Left = &node1
		node1.Left = &node2
		node1.Right = &node3
		node3.Left = &node4
		node3.Right = &node5

		root.Right = &node6
		node6.Left = &node7
		node6.Right = &node8

		p := TreeNode{Val: 5}
		q := TreeNode{Val: 1}
		expect := TreeNode{Val: 3}
		got := LowestCommonAncestor1(&root, &p, &q)
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}

		q = TreeNode{Val: 5}
		q = TreeNode{Val: 4}
		expect = TreeNode{Val: 5}
		got = LowestCommonAncestor1(&root, &p, &q)
		if got.Val != expect.Val {
			t.Fatalf("expect:%v, got:%v", expect.Val, got.Val)
		}
	})
}
