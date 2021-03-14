package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var successor *ListNode
	if left == 1 {
		return reverseN(head, left, successor)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverseN(head *ListNode, n int, successor *ListNode) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1, successor)
	head.Next.Next = head
	head.Next = successor
	return last
}
