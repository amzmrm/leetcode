package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	reverse(head, &res)
	return res
}

func reverse(head *ListNode, res *[]int) {
	if head == nil {
		return
	}

	if head.Next != nil {
		reverse(head.Next, res)
	}
	*res = append(*res, head.Val)

	return
}
