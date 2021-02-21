package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	idx := indexOfMaxNum(nums)

	root := TreeNode{Val: nums[idx]}
	left := constructMaximumBinaryTree(nums[:idx])
	right := constructMaximumBinaryTree(nums[idx+1:])

	root.Left, root.Right = left, right

	return &root
}

func indexOfMaxNum(nums []int) int {
	max := nums[0]
	idx := 0
	for i, num := range nums {
		if num > max {
			max = num
			idx = i
		}
	}
	return idx
}
