package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var numbers []int
	inorder(root, &numbers)
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == k {
			return true
		} else if sum > k {
			right--
		} else if sum < k {
			left++
		}
	}
	return false
}

func inorder(root *TreeNode, numbers *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, numbers)
	*numbers = append(*numbers, root.Val)
	inorder(root.Right, numbers)
}
