package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil { // 比较相等需要两个数，当前节点的下一个和下下一个
		if curr.Next.Val == curr.Next.Next.Val {
			val := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == val { // 每次删除一个，而不是一下子把相同的那段删除
				curr.Next = curr.Next.Next // 注意：这里curr没有向前移动，也不要移动，而是找到重复区间后的第一个节点
			}
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}
