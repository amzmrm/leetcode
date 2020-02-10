package problem_112_path_sum

import "testing"

func TestHasPathSum(t *testing.T) {
	type test struct {
		name   string
		tree   *TreeNode
		sum    int
		expect bool
	}

	tests := make([]test, 0)

	normalCase := &TreeNode{Val:5}
	t1n2 := &TreeNode{Val:4}
	t1n3 := &TreeNode{Val:8}
	t1n4 := &TreeNode{Val:11}
	t1n5 := &TreeNode{Val:7}
	t1n6 := &TreeNode{Val:2}
	t1n7 := &TreeNode{Val:13}
	t1n8 := &TreeNode{Val:4}
	t1n9 := &TreeNode{Val:1}
	normalCase.Left = t1n2
	normalCase.Right = t1n3
	t1n2.Left = t1n4
	t1n4.Left = t1n5
	t1n4.Right = t1n6
	t1n3.Left = t1n7
	t1n3.Right = t1n8
	t1n8.Right = t1n9
	tests = append(tests, test{
		name:   "normal case",
		tree:   normalCase,
		sum:    22,
		expect: true,
	})

	var emptyTree *TreeNode
	tests = append(tests, test{
		name:   "empty tree",
		tree:   emptyTree,
		sum:    22,
		expect: false,
	})

	singleNodeTree := &TreeNode{Val:22}
	tests = append(tests, test{
		name:   "single node tree",
		tree:   singleNodeTree,
		sum:    22,
		expect: true,
	})

	negative := &TreeNode{Val: 5}
	t4n2 := &TreeNode{Val:3}
	t4n3 := &TreeNode{Val:8}
	t4n4 := &TreeNode{Val:11}
	t4n5 := &TreeNode{Val:7}
	t4n6 := &TreeNode{Val:2}
	t4n7 := &TreeNode{Val:13}
	t4n8 := &TreeNode{Val:4}
	t4n9 := &TreeNode{Val:1}
	negative.Left = t4n2
	negative.Right = t4n3
	t4n2.Left = t4n4
	t4n4.Left = t4n5
	t4n4.Right = t4n6
	t4n3.Left = t4n7
	t4n3.Right = t4n8
	t4n8.Right = t4n9
	tests = append(tests, test{
		name:   "negative",
		tree:   negative,
		sum:    22,
		expect: false,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := HasPathSumByIterative(test.tree, test.sum)
			if test.expect != got {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
