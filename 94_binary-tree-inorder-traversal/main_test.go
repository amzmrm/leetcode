package leetcode

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	type test struct {
		name   string
		tree   *TreeNode
		expect []int
	}

	tests := make([]test, 0)

	// 正常情况
	root1Node1 := &TreeNode{Val: 1}
	root1Node2 := &TreeNode{Val: 2}
	root1Node3 := &TreeNode{Val: 3}
	root1Node4 := &TreeNode{Val: 4}
	root1Node5 := &TreeNode{Val: 5}
	root1Node6 := &TreeNode{Val: 6}
	root1Node1.Left = root1Node2
	root1Node1.Right = root1Node3
	root1Node2.Left = root1Node4
	root1Node2.Right = root1Node5
	root1Node3.Right = root1Node6
	tests = append(tests, test{
		name:   "normal",
		tree:   root1Node1,
		expect: []int{4, 2, 5, 1, 3, 6},
	})

	// 空二叉树
	var root2Node1 *TreeNode
	tests = append(tests, test{
		name:   "empty tree",
		tree:   root2Node1,
		expect: []int{},
	})

	// 只有根节点二叉树
	root3Node1 := &TreeNode{Val: 1}
	tests = append(tests, test{
		name:   "root node only",
		tree:   root3Node1,
		expect: []int{1},
	})

	// 只有左子树
	root4Node1 := &TreeNode{Val: 1}
	root4Node2 := &TreeNode{Val: 2}
	root4Node1.Left = root4Node2
	tests = append(tests, test{
		name:   "root and left child only",
		tree:   root4Node1,
		expect: []int{2, 1},
	})

	// 只有右子树
	root5Node1 := &TreeNode{Val: 1}
	root5Node2 := &TreeNode{Val: 2}
	root5Node1.Right = root5Node2
	tests = append(tests, test{
		name:   "root and right child only",
		tree:   root5Node1,
		expect: []int{1, 2},
	})

	// leetcode测试例子
	root6Node1 := &TreeNode{Val: 1}
	root6Node2 := &TreeNode{Val: 2}
	root6Node3 := &TreeNode{Val: 3}
	root6Node1.Right = root6Node2
	root6Node2.Left = root6Node3
	tests = append(tests, test{
		name:   "leetcode",
		tree:   root6Node1,
		expect: []int{1, 3, 2},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MorrisInorderTraversal(test.tree)
			if !reflect.DeepEqual(got, test.expect) {
				t.Fatalf("got:%#v, expect:%#v", got, test.expect)
			}
		})
	}
}
