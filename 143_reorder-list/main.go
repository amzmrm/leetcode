package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 1. 先使用快慢指针找到链表中点
	mid := middleNode(head)

	// 2. 反转中点后半部分
	l1, l2 := head, mid.Next
	mid.Next = nil
	l2 = reverse(l2)

	// 3. 开始合并两个链表
	var l1Next, l2Next *ListNode
	for l1 != nil && l2 != nil {
		// 暂存两个链表表头的下一个指针
		l1Next = l1.Next
		l2Next = l2.Next

		l1.Next = l2     // l1.Next指向另一个表头
		l2.Next = l1Next // l2.Next应该指向l1的下一个开始的位置
		l1 = l1Next      // l1表头向下一位移动一步
		l2 = l2Next      // l2表头向下一位移动一步
	}
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
	var prev, next *ListNode
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
