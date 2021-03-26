package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateList(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var length int
	var tail *ListNode
	for curr := head; curr != nil; {
		tail = curr
		length++
		curr = curr.Next
	}
	steps := k % length
	if steps == 0 {
		return head
	}

	newTail := getKthFromEnd(head, steps+1)
	newHead := newTail.Next
	tail.Next = head
	newTail.Next = nil
	return newHead
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
