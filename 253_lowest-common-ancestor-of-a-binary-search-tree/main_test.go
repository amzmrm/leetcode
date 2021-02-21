package leetcode

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	type test struct {
		name   string
		tree   *TreeNode
		p, q   *TreeNode
		expect int
	}

	tests := make([]test, 0)

	t1n1 := &TreeNode{Val: 6}
	t1n2 := &TreeNode{Val: 2}
	t1n3 := &TreeNode{Val: 8}
	t1n4 := &TreeNode{Val: 0}
	t1n5 := &TreeNode{Val: 4}
	t1n6 := &TreeNode{Val: 7}
	t1n7 := &TreeNode{Val: 9}
	t1n8 := &TreeNode{Val: 3}
	t1n9 := &TreeNode{Val: 5}

	t1n1.Left = t1n2
	t1n2.Left = t1n4
	t1n2.Right = t1n5
	t1n5.Left = t1n8
	t1n5.Right = t1n9
	t1n1.Right = t1n3
	t1n3.Left = t1n6
	t1n3.Right = t1n7

	tests = append(tests, test{
		name:   "normal1",
		tree:   t1n1,
		p:      t1n2,
		q:      t1n3,
		expect: 6,
	})
	tests = append(tests, test{
		name:   "normal2",
		tree:   t1n1,
		p:      t1n5,
		q:      t1n9,
		expect: 4,
	})

	for _, test := range tests {
		got := LowestCommonAncestor(test.tree, test.p, test.q)
		if got.Val != test.expect {
			t.Fatalf("got:%v, expect:%v", got.Val, test.expect)
		}
	}
}
