package problem_25_reverse_nodes_in_k_group

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
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil { // 空链表直接返回
		return head
	}
	if k <= 1 { // k==1表示不翻转
		return head
	}

	// 先定位到新的链表头
	// 新的链表头是第一组翻转后的链表头
	cursor := 0
	newHead := head
	for cursor < k-1 && newHead.Next != nil {
		newHead = newHead.Next
		cursor++
	}
	if cursor < k-1 { // 链表长度小于k，直接返回原链表
		return head
	}

	var next *ListNode      // 用于保存当前分组的下一个链表节点
	var groupTail *ListNode // 反转前一个组的尾节点
	var preHead *ListNode   // 上一组反转后的尾节点，下一个组反转后将上一组的尾节点接上下一组的头节点
	var groupHead *ListNode // 反转前一个组的头节点
	cursor = 0              // 用于帮助遍历得到一个组
	curr := head            // 当前遍历到的节点
	for curr != nil {
		if cursor == 0 {
			groupHead = curr
			cursor++
			continue
		} else if cursor < k {
			curr = curr.Next
			cursor++
			continue
		}
		groupTail = curr
		curr = curr.Next

		// 开始做翻转
		next = groupTail.Next // 保存当前分组的下一个链表节点

		var groupPrehead, groupNext *ListNode
		groupCurr := groupHead
		for groupCurr != next { // 翻转这一组的链表，其实就是单链表翻转
			groupNext = groupCurr.Next
			groupCurr.Next = groupPrehead
			groupPrehead = groupCurr
			groupCurr = groupNext
		}

		// 一组翻转结束后
		if preHead != nil {
			preHead.Next = groupTail // 上一组反转前的尾接上本组反转前的尾
		}
		preHead = groupHead

		cursor = 0 // 完成一个组的反转，将cursor置为0，进行下一组的反转
	}

	// 接上非k整数倍的剩余节点
	// 非整数倍的情况，根据上面得知cursor不会置为0
	if cursor > 0 {
		preHead.Next = groupHead
	}

	return newHead
}
