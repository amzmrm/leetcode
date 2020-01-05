package problem_94_binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归公式：inorderTraversal(node) = inorderTraversal(node.Left) + node.Val + inorderTraversal(node.Right)
// 递归终止条件：node为nil
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := append(InorderTraversal(root.Left), root.Val)
	ret = append(ret, InorderTraversal(root.Right)...)

	return ret
}
