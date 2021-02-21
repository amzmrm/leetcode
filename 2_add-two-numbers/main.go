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

// 关键是要使用两个指针，preHead用来指向新链表的头节点，prev用来在遍历链表的过程中将各个新节点串联起来
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2

	preHead := &ListNode{}
	prev := preHead

	carry := 0
	for p1 != nil || p2 != nil {
		val1, val2 := 0, 0
		if p1 != nil {
			val1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val2 = p2.Val
			p2 = p2.Next
		}

		newNode := &ListNode{}
		newNode.Val = val1 + val2 + carry
		carry = 0
		if newNode.Val >= 10 {
			newNode.Val = newNode.Val % 10
			carry = 1
		}
		prev.Next = newNode
		prev = newNode
	}

	if carry > 0 {
		newNode := &ListNode{}
		newNode.Val = carry
		carry = 0

		prev.Next = newNode
		prev = newNode
	}

	return preHead.Next
}
