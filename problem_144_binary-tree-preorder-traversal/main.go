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

func IterativePreorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	stack := make([]*TreeNode, 0)

	curr := root
	for curr != nil || len(stack) > 0 {
		if curr == nil {
			curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
			continue
		}

		values = append(values, curr.Val)
		if curr.Left != nil {
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return values
}