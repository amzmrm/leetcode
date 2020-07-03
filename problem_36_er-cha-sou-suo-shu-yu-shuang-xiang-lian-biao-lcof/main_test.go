package main

import (
	"reflect"
	"testing"
)

func TestConvertToDoubleList(t *testing.T) {
	t.Parallel()

	t.Run("normal", func(t *testing.T) {
		root := TreeNode{Val: 4}
		node2 := TreeNode{Val: 2}
		node3 := TreeNode{Val: 1}
		node4 := TreeNode{Val: 3}
		node5 := TreeNode{Val: 5}
		root.Left = &node2
		root.Right = &node5
		node2.Left = &node3
		node2.Right = &node4

		result := ConvertToDoubleList(&root)

		if !reflect.DeepEqual(result.Traversal(), []int{1, 2, 3, 4, 5}) {
			t.Fatalf("got:%v, expect:%v", result.Traversal(), []int{1, 2, 3, 4, 5})
		}
		if !reflect.DeepEqual(result.ReverseTraversal(), []int{5, 4, 3, 2, 1}) {
			t.Fatalf("got:%v, expect:%v", result.ReverseTraversal(), []int{5, 4, 3, 2, 1})
		}
	})

	t.Run("single node tree", func(t *testing.T) {
		root := TreeNode{Val: 4}
		result := ConvertToDoubleList(&root)

		if !reflect.DeepEqual(result.Traversal(), []int{4}) {
			t.Fatalf("got:%v, expect:%v", result.Traversal(), []int{4})
		}

		if !reflect.DeepEqual(result.ReverseTraversal(), []int{4}) {
			t.Fatalf("got:%v, expect:%v", result.ReverseTraversal(), []int{4})
		}
	})

	t.Run("two nodes tree", func(t *testing.T) {
		root := TreeNode{Val: 4}
		node2 := TreeNode{Val: 2}
		root.Left = &node2
		result := ConvertToDoubleList(&root)

		expect := []int{2, 4}
		got := result.Traversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}

		expect = []int{4, 2}
		got = result.ReverseTraversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}

		root = TreeNode{Val: 4}
		node2 = TreeNode{Val: 5}
		root.Right = &node2
		result = ConvertToDoubleList(&root)
		expect = []int{4, 5}
		got = result.Traversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}

		expect = []int{5, 4}
		got = result.ReverseTraversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}
	})

	t.Run("full binary tree", func(t *testing.T) {
		root := TreeNode{Val: 2}
		node2 := TreeNode{Val: 1}
		node3 := TreeNode{Val: 3}
		root.Left = &node2
		root.Right = &node3

		result := ConvertToDoubleList(&root)

		expect := []int{1, 2, 3}
		got := result.Traversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}

		expect = []int{3, 2, 1}
		got = result.ReverseTraversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}
	})

	t.Run("child only on the left", func(t *testing.T) {
		root := TreeNode{Val: 5}
		node2 := TreeNode{Val: 4}
		node3 := TreeNode{Val: 3}
		node4 := TreeNode{Val: 2}
		node5 := TreeNode{Val: 1}
		root.Left = &node2
		node2.Left = &node3
		node3.Left = &node4
		node4.Left = &node5

		result := ConvertToDoubleList(&root)

		expect := []int{1, 2, 3, 4, 5}
		got := result.Traversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}

		expect = []int{5, 4, 3, 2, 1}
		got = result.ReverseTraversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}
	})

	t.Run("child only on the right", func(t *testing.T) {
		root := TreeNode{Val: 1}
		node2 := TreeNode{Val: 2}
		node3 := TreeNode{Val: 3}
		node4 := TreeNode{Val: 4}
		node5 := TreeNode{Val: 5}
		root.Right = &node2
		node2.Right = &node3
		node3.Right = &node4
		node4.Right = &node5

		result := ConvertToDoubleList(&root)

		expect := []int{1, 2, 3, 4, 5}
		got := result.Traversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}

		expect = []int{5, 4, 3, 2, 1}
		got = result.ReverseTraversal()
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("got:%v, expect:%v", got, expect)
		}
	})

	t.Run("nil tree", func(t *testing.T) {
		var root *TreeNode
		result := ConvertToDoubleList(root)
		if !reflect.DeepEqual(root, result) {
			t.Fatalf("got:%v, expect:%v", result, root)
		}
	})
}
