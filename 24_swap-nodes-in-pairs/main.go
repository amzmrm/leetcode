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

// 注意：代码逻辑在处理头结点和尾结点的时候，是否能正常工作
func SwapPairs(head *ListNode) *ListNode {
	if head == nil { // 空链表检测
		return nil
	}
	if head.Next == nil { // 只有一个节点的链表检测
		return head
	}

	curr, next := head, head.Next
	preHead := next // 新的链表头节点就是第一对的开始节点

	lastPairEnd := curr         // 上一对的结束
	var nextPairStart *ListNode // 下一对的开始

	for curr != nil && next != nil {
		nextPairStart = next.Next

		lastPairEnd.Next = next
		next.Next = curr
		curr.Next = nextPairStart
		lastPairEnd = curr

		curr = nextPairStart
		if nextPairStart != nil {
			next = nextPairStart.Next
		} else {
			next = nil
		}
	}

	return preHead
}
