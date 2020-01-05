package problem_144_binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归公式：preorderTraversal(node) = node.Val + preorderTraversal(node.Left) + preorderTraversal(node.Right)
// 递归终止条件：node为nil
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := append([]int{root.Val}, PreorderTraversal(root.Left)...)
	ret = append(ret, PreorderTraversal(root.Right)...)

	return ret
}

