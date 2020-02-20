package problem_141_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
