package problem_07_zhong_jian_er_cha_shu_lcof

// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 例如，给出
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：
//    3
//   / \
//  9  20
//    /  \
//   15   7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	if len(preorder) == 1 || len(inorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	return build(preorder, inorder)
}

func build(preorder []int, inorder []int) *TreeNode {
	rootVal := preorder[0]
	rootIndex := rootIndex(inorder, rootVal)
	leftTreeLen, rightTreeLen := subTreeLen(inorder, rootIndex)

	root := &TreeNode{Val: rootVal}

	if leftTreeLen > 1 {
		root.Left = build(preorder[1:1+leftTreeLen], inorder[:rootIndex])
	} else if leftTreeLen == 1 {
		root.Left = &TreeNode{Val: inorder[0]}
	} else {
		// no left sub tree
	}

	if rightTreeLen > 0 {
		root.Right = build(preorder[1+leftTreeLen:], inorder[rootIndex+1:])
	} else if rightTreeLen == 1 {
		root.Right = &TreeNode{Val: inorder[rootIndex+1]}
	} else {
		// no right sub tree
	}

	return root
}

func rootIndex(inorder []int, rootVal int) int {
	for i, num := range inorder {
		if num == rootVal {
			return i
		}
	}
	return -1
}

func subTreeLen(inorder []int, i int) (int, int) {
	left := len(inorder[:i])
	right := 0
	if i < len(inorder) {
		right = len(inorder[i+1:])
	}
	return left, right
}

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := append([]int{root.Val}, PreorderTraversal(root.Left)...)
	ret = append(ret, PreorderTraversal(root.Right)...)

	return ret
}
