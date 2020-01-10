package problem_144_binary_tree_preorder_traversal

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
	root1Node3.Left = root1Node6
	tests = append(tests, test{
		name:   "normal",
		tree:   root1Node1,
		expect: []int{1, 2, 4, 5, 3, 6},
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
		expect: []int{1, 2},
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
		expect: []int{1, 2, 3},
	})

	// 迭代遍历测试用例
	root7Node1 := &TreeNode{Val: 1}
	root7Node2 := &TreeNode{Val: 2}
	root7Node3 := &TreeNode{Val: 3}
	root7Node4 := &TreeNode{Val: 4}
	root7Node5 := &TreeNode{Val: 5}
	root7Node6 := &TreeNode{Val: 6}
	root7Node1.Right = root7Node2
	root7Node2.Right = root7Node3
	root7Node3.Right = root7Node4
	root7Node4.Left = root7Node5
	root7Node5.Right = root7Node6
	tests = append(tests, test{
		name:   "Iterative traversal",
		tree:   root7Node1,
		expect: []int{1, 2, 3, 4, 5, 6},
	})

	// Morris遍历测试用例
	root8Node1 := &TreeNode{Val:1}
	root8Node2 := &TreeNode{Val:2}
	root8Node3 := &TreeNode{Val:3}
	root8Node4 := &TreeNode{Val:4}
	root8Node5 := &TreeNode{Val:5}
	root8Node1.Left = root8Node2
	root8Node1.Right = root8Node5
	root8Node2.Left = root8Node3
	root8Node2.Right = root8Node4
	tests = append(tests, test{
		name:   "Morris traversal",
		tree:   root8Node1,
		expect: []int{1, 2, 3, 4, 5},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MorrisPreorderTraversal(test.tree)
			if !reflect.DeepEqual(PreorderTraversal(test.tree), test.expect) {
				t.Fatalf("got:%#v, expect:%#v", got, test.expect)
			}
		})
	}
}
