package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func LowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := LowestCommonAncestor1(root.Left, p, q)
	right := LowestCommonAncestor1(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parents := make(map[int]*TreeNode)
	visited := make(map[int]bool)

	var dfs func(node *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		if r.Left != nil {
			parents[r.Left.Val] = r
			dfs(r.Left)
		}

		if r.Right != nil {
			parents[r.Right.Val] = r
			dfs(r.Right)
		}
	}

	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parents[p.Val]
	}

	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parents[q.Val]
	}

	return nil
}
