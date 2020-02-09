package problem_104_maximum_depth_of_binary_tree

import "testing"

func TestMaxDepth(t *testing.T) {
	type test struct {
		name   string
		tree   *TreeNode
		expect int
	}

	tests := make([]test, 0)

	t1n1 := &TreeNode{Val: 3}
	t1n2 := &TreeNode{Val: 9}
	t1n3 := &TreeNode{Val: 20}
	t1n4 := &TreeNode{Val: 15}
	t1n5 := &TreeNode{Val: 7}
	t1n1.Left = t1n2
	t1n1.Right = t1n3
	t1n3.Left = t1n4
	t1n3.Right = t1n5
	tests = append(tests, test{
		name:   "leetcode example",
		tree:   t1n1,
		expect: 3,
	})

	var emptyTree *TreeNode
	tests = append(tests, test{
		name:   "empty_tree",
		tree:   emptyTree,
		expect: 0,
	})

	singleNodeTree := &TreeNode{Val:3}
	tests = append(tests, test{
		name:   "single node tree",
		tree:   singleNodeTree,
		expect: 1,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MaxDepthIterative(test.tree)
			if got != test.expect {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
