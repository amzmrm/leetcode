package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	// 如果第一次遍历到链表尾部，就指向另一个链表的头部，继续遍历，这样会抵消长度差。如果没有相交，因为遍历长度相等，最后会是nil == nil
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}
