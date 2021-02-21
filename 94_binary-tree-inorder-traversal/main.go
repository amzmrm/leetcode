package leetcode

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

func IterativeInorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}

	// 使用slice模拟一个栈
	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) > 0 {
		// 循环遍历左子树，直到叶子节点
		// 只要curr不是nil，不管是左子树还是右子树都入栈
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// 左子树遍历遇到了叶子节点，开始出栈
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
		vals = append(vals, curr.Val) // 子树的根节点
		curr = curr.Right             /// 开始遍历右子树
	}

	return vals
}

func MorrisInorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}

	curr := root
	var pre *TreeNode
	for curr != nil {
		if curr.Left == nil {
			vals = append(vals, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr
			curr, curr.Left = curr.Left, nil
		}
	}
	return vals
}
