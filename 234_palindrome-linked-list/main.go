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

// 1. 首先找到前半部分的尾节点
// 2. 反转后半部分
// 3. 前后部分对比
// 4. 复原后半部分
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	middle := middleNode(head)
	secondHalfStart := reverse(middle.Next)

	firstPos, secondPos := head, secondHalfStart
	for secondPos != nil {
		if firstPos.Val != secondPos.Val {
			return false
		}
		firstPos = firstPos.Next
		secondPos = secondPos.Next
	}

	middle.Next = reverse(secondHalfStart)

	return true
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
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
