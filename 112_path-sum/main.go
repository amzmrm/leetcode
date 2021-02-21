package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归公式：hasPathSum(root,sum) = hasPathSum(root.Left,sum-root.Val) || hasPathSum(root.Right,sum-root.Val)
// 递归推出条件：(sum - root.Val == 0) || (root.Left == nil && root.Right == nil)为true
// 时间复杂度：O(n)
// 空间复杂度：如果树完全平衡O(logn)，如果树每个节点都只有一个节点O(n)
func HasPathSumByRecursively(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if sum == 0 {
		return true
	}

	// 如果当前节点是叶子结点，此次函数调用返回
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return HasPathSumByRecursively(root.Left, sum) || HasPathSumByRecursively(root.Right, sum)
}

func HasPathSumByIterative(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	type element struct {
		node *TreeNode
		sum  int
	}

	stack := make([]*element, 0)
	stack = append(stack, &element{
		node: root,
		sum:  sum - root.Val,
	})

	var elem *element
	for len(stack) > 0 {
		elem, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if elem.node.Left == nil && elem.node.Right == nil && elem.sum == 0 {
			return elem.sum == 0
		}

		if elem.node.Right != nil {
			stack = append(stack, &element{
				node: elem.node.Right,
				sum:  elem.sum - elem.node.Right.Val,
			})
		}
		if elem.node.Left != nil {
			stack = append(stack, &element{
				node: elem.node.Left,
				sum:  elem.sum - elem.node.Left.Val,
			})
		}
	}

	return false
}
