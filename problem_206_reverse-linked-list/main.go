package problem_206_reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToSlice() []int {
	nums := make([]int, 0)

	curr := l
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}
	return nums
}

func ReverseList(head *ListNode) *ListNode {
	curr := head
	var preHead, next *ListNode
	for curr != nil {
		next = curr.Next
		curr.Next = preHead
		preHead = curr
		curr = next
	}
	return preHead
}
