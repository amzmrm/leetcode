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
	if root == nil {
		return values
	}

	stack := make([]*TreeNode, 0)
	var node *TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		values = append(values, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return values
}

// 时间复杂度：O(n)
// 空间复杂度：只遍历输出时是O(1)，遍历并返回时是O(n)
// Morris算法避免了使用栈。栈的作用是开始遍历root的左子树时，需要暂存root的右子树节点，因为遍历完左子树后要接着遍历右子树，需要用到右子树的指针
// Morris算法让root的左子树的最右边叶子节点的右指针指向了root，因此遍历完左子树后能再次拿到root指针，从而也就能拿到root的右子树的指针
func MorrisPreorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)

	var predecessor *TreeNode
	curr := root
	for curr != nil {
		if curr.Left == nil {
			// 左子树为空，输出当前根节点后，直接遍历到右子树
			// 这里的右子树可能是真的右子树，也可能是伪右子树
			values = append(values, curr.Val)
			curr = curr.Right
			continue
		}

		// 左子树不为空，找到左子树最右边的节点
		predecessor = curr.Left // predecessor要么是它最右边的叶子节点，要么就是它自己
		for predecessor.Right != nil && predecessor.Right != curr {
			predecessor = predecessor.Right
		}

		if predecessor.Right == nil {
			// 第一次从curr遍历到该predecessor.Right时，肯定是nil，因为树无论如何都会有叶子节点
			// 遍历到curr的最右边叶子节点后，输出curr节点的值，并将最右边叶子节点的右指针指向curr
			// 最后将curr往左子树方向下移一步
			values = append(values, curr.Val)
			predecessor.Right = curr
			curr = curr.Left
		} else {
			// 第二次从curr遍历到predecessor.Right时，predecessor.Right早已指向对应的curr，不为nil
			// 第二次从同一个curr遍历下来，这时要将curr往右子树方向走一步
			// 并且需要将这个节点的伪右子树的指针置为nil
			curr = curr.Right
			predecessor.Right = nil
		}
	}

	return values
}
