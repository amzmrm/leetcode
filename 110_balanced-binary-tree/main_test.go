package leetcode

import "testing"

func TestIsBalanced(t *testing.T) {
	type test struct {
		name   string
		tree   *TreeNode
		expect bool
	}

	tests := make([]test, 0)

	var t1n1 *TreeNode
	tests = append(tests, test{
		name:   "empty tree",
		tree:   t1n1,
		expect: true,
	})

	t2n1 := &TreeNode{Val: 1}
	tests = append(tests, test{
		name:   "root only tree",
		tree:   t2n1,
		expect: true,
	})

	t3n1 := &TreeNode{Val: 3}
	t3n2 := &TreeNode{Val: 9}
	t3n3 := &TreeNode{Val: 20}
	t3n4 := &TreeNode{Val: 15}
	t3n5 := &TreeNode{Val: 7}
	t3n1.Left = t3n2
	t3n1.Right = t3n3
	t3n3.Left = t3n4
	t3n3.Right = t3n5
	tests = append(tests, test{
		name:   "leetcode case 1",
		tree:   t3n1,
		expect: true,
	})

	t4n1 := &TreeNode{Val: 1}
	t4n2 := &TreeNode{Val: 2}
	t4n3 := &TreeNode{Val: 3}
	t4n4 := &TreeNode{Val: 3}
	t4n5 := &TreeNode{Val: 4}
	t4n6 := &TreeNode{Val: 4}
	t4n7 := &TreeNode{Val: 2}
	t4n1.Left = t4n2
	t4n2.Right = t4n7
	t4n2.Left = t4n3
	t4n2.Right = t4n4
	t4n3.Left = t4n5
	t4n3.Right = t4n6
	tests = append(tests, test{
		name:   "leetcode case 2",
		tree:   t4n1,
		expect: false,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsBalanced(test.tree)
			if got != test.expect {
				t.Fatalf("got: %#v, expect:%#v", got, test.expect)
			}
		})
	}
}
