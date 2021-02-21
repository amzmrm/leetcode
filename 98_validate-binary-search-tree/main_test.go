package leetcode

import "testing"

func TestIsValidBST(t *testing.T) {
	type test struct {
		name   string
		tree   *TreeNode
		expect bool
	}

	tests := make([]test, 0)

	t1n1 := &TreeNode{Val: 2}
	t1n2 := &TreeNode{Val: 1}
	t1n3 := &TreeNode{Val: 3}
	t1n1.Left = t1n2
	t1n1.Right = t1n3
	tests = append(tests, test{
		name:   "simple binary search tree",
		tree:   t1n1,
		expect: true,
	})

	t2n1 := &TreeNode{Val: 5}
	t2n2 := &TreeNode{Val: 1}
	t2n3 := &TreeNode{Val: 4}
	t2n4 := &TreeNode{Val: 3}
	t2n5 := &TreeNode{Val: 6}
	t2n1.Left = t2n2
	t2n1.Right = t2n3
	t2n3.Left = t2n4
	t2n4.Right = t2n5
	tests = append(tests, test{
		name:   "non binary search tree",
		tree:   t2n1,
		expect: false,
	})

	var emptyTree *TreeNode
	tests = append(tests, test{
		name:   "empty tree",
		tree:   emptyTree,
		expect: true,
	})

	singleNodeTree := &TreeNode{Val: 0}
	tests = append(tests, test{
		name:   "single node tree",
		tree:   singleNodeTree,
		expect: true,
	})

	cornerCase1 := &TreeNode{Val: 5}
	t4n2 := &TreeNode{Val: 1}
	t4n3 := &TreeNode{Val: 6}
	t4n4 := &TreeNode{Val: 4}
	t4n5 := &TreeNode{Val: 7}
	cornerCase1.Left = t4n2
	cornerCase1.Right = t4n3
	t4n3.Left = t4n4
	t4n4.Right = t4n5
	tests = append(tests, test{
		name:   "corner case 1",
		tree:   cornerCase1,
		expect: false,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsValidBST(test.tree)
			if got != test.expect {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
