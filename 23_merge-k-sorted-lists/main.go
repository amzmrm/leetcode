package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type head struct {
	node  *ListNode
	index int
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

// 时间复杂度：O(nlogk)
// 空间复杂度：O(1)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	amount := len(lists)
	interval := 1
	for interval < amount {
		for i := 0; i < amount-interval; i += 2 * interval {
			lists[i] = merge2Lists(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
}

// 时间复杂度：O(m+n)，其中m,n分别为两个链表的长度
// 空间复杂度：O(1)
func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	prev := preHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}

	return preHead.Next
}
