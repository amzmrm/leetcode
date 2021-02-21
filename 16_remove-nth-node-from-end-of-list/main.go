package leetcode

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

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	first, second := head, head

	var i int
	for i = 0; i < n+1 && first != nil; i++ {
		first = first.Next
	}

	if i < n+1 {
		// 移除倒数第len(linkedList)位，first指针和second指针之间的差距比len还大
		return head.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return head
}
