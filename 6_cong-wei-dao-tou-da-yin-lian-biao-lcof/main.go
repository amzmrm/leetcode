package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func ReversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	vals := make([]int, 0)
	curr := head
	for curr != nil {
		vals = append([]int{curr.Val}, vals...)
		curr = curr.Next
	}

	return vals
}
