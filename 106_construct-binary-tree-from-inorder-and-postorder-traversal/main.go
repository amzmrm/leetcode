package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := TreeNode{Val: postorder[len(postorder)-1]}
	var idx int
	for i, node := range inorder {
		if node == root.Val {
			idx = i
			break
		}
	}

	// 知道root后，就知道根结点的左子树和右子树了
	inorderLeft, inorderRight := inorder[:idx], inorder[idx+1:]
	// 知道根结点的左子树和右子树后，就知道左右子树的长度了
	postorderLeftLen, postorderRightLen := len(inorderLeft), len(inorderRight)
	// 知道右子树的长度后，postorder中右子树的起点就是最后一个节点往回数postorderRightLen个
	rStart := len(postorder) - 1 - postorderRightLen
	postorderRight := postorder[rStart : len(postorder)-1]
	// 知道左子树的长度后，postorder中左子树的起点就是右子树的起点往回数postorderLeftLen个
	lStart := rStart - postorderLeftLen
	postorderLeft := postorder[lStart:rStart]

	root.Left = buildTree(inorderLeft, postorderLeft)
	root.Right = buildTree(inorderRight, postorderRight)

	return &root
}
