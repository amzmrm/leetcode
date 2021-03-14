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

func ReverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	for head.Next != nil {
		t := head.Next.Next
		head.Next.Next = curr
		curr = head.Next
		head.Next = t
	}
	return curr
}

func ReverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := ReverseListRecursively(head.Next) // 一直返回尾节点，也就是递归结束条件那个节点

	// 局部反转
	head.Next.Next = head // 反转2个节点之间的指针
	head.Next = nil

	return ret
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
