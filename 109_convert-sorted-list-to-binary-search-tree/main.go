package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return buildTree(head, nil)
}

func buildTree(start *ListNode, end *ListNode) *TreeNode {
	if start == end {
		return nil
	}
	mid := middleNode(start, end)
	root := &TreeNode{Val: mid.Val}
	root.Left = buildTree(start, mid)
	root.Right = buildTree(mid.Next, end)
	return root
}

func middleNode(start *ListNode, end *ListNode) *ListNode {
	slow, fast := start, start
	for fast != end && fast.Next != end {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
