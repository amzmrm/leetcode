package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := LowestCommonAncestor1(root.Left, p, q)   // returns nil or p or q
	right := LowestCommonAncestor1(root.Right, p, q) // returns nil or p or q
	if left != nil && right != nil {
		// this node's left child and right child both are not nil, this node is the lowest common ancestor
		return root
	}
	if left == nil {
		// right could also be nil
		return right
	}
	return left // could nil or not nil
}

func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parents := make(map[int]*TreeNode)
	visited := make(map[int]bool)

	// traversal over the tree and save each element's parent node
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

	// visit all p's ancestor and mark as visited
	for p != nil {
		visited[p.Val] = true
		p = parents[p.Val]
	}

	// visit all q's ancestor and return the one was previously visited
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parents[q.Val]
	}

	return nil
}
