package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type limit struct {
	val int
}

func IsValidBSTByRecursively(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return helper(root, nil, nil)
}

func helper(node *TreeNode, lower, upper *limit) bool {
	if node == nil {
		return true
	}

	if lower != nil && node.Val <= lower.val {
		return false
	}
	if upper != nil && node.Val >= upper.val {
		return false
	}

	if !helper(node.Left, lower, &limit{val: node.Val}) {
		return false
	}
	if !helper(node.Right, &limit{val: node.Val}, upper) {
		return false
	}

	return true
}

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 使用slice模拟一个栈
	stack := make([]*TreeNode, 0)

	curr := root
	inorder := math.MinInt64
	for curr != nil || len(stack) > 0 {
		// 只要curr不是nil，不管是左子树还是右子树都入栈
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// 左子树遍历遇到了叶子节点，开始出栈
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]

		// 遍历序列越往前越大，如果不是则返回false
		if curr.Val <= inorder {
			return false
		}

		inorder = curr.Val
		curr = curr.Right // 开始遍历右子树
	}

	return true
}

func IsIsValidBSTV2(root *TreeNode) bool {
	return IsValidBSTHelpV2(root, nil, nil)
}

func IsValidBSTHelpV2(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return IsValidBSTHelpV2(root.Left, min, root) && IsValidBSTHelpV2(root.Right, root, max)
}
