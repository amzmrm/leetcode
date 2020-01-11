package problem_145_binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IterativePostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	values := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	var node *TreeNode
	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		values = append([]int{node.Val}, values...)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return values
}
