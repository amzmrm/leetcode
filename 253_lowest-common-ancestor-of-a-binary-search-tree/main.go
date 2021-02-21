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

	return helper2(root, min, max)
}

// 递归找出第一个值在[min,max)之间的节点
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func helper1(root *TreeNode, min int, max int) *TreeNode {
	if max < root.Val {
		return helper1(root.Left, min, max)
	}

	if root.Val < min {
		return helper1(root.Right, min, max)
	}

	return root
}

// 迭代找出第一个值在[min,max)之间的节点
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func helper2(root *TreeNode, min int, max int) *TreeNode {
	node := root
	for node != nil {
		if max < node.Val {
			node = node.Left
			continue
		}
		if node.Val < min {
			node = node.Right
			continue
		}
		return node
	}
	return nil
}
