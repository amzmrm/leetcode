package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	if len(nums[:mid]) > 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	if len(nums[mid:])-1 > 0 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}
