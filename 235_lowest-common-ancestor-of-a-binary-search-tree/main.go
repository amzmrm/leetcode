package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	min, max := p.Val, q.Val
	if q.Val < min {
		min, max = max, min
	}

	return helper(root, min, max)
}

// 递归找出第一个值在[min,max)之间的节点
func helper(root *TreeNode, min int, max int) *TreeNode {
	if max < root.Val {
		return helper(root.Left, min, max)
	}

	if root.Val < min {
		return helper(root.Right, min, max)
	}

	return root
}
