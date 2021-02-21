package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 事件复杂度：O(n)
// 空间复杂度：O(n)
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//node, _ := helper1(root, p, q)
	//return node

	return helper2(root, p, q)
}

func helper1(currentNode, p, q *TreeNode) (*TreeNode, bool) {
	if currentNode == nil {
		return nil, false
	}

	// left表示向左子树递归是否能找到通向p/q，能找到赋值1，不能的话赋值0
	// right表示向右子树递归是否能找到通向p/q，能找到赋值1，不能的话赋值0
	// mid表示当前节点是否等于p/q，等于的话赋值1，不等于赋值0
	mid, left, right := 0, 0, 0

	node, ok1 := helper1(currentNode.Left, p, q)
	if node != nil {
		return node, true
	}
	if ok1 {
		left = 1
	}

	node, ok2 := helper1(currentNode.Right, p, q)
	if node != nil {
		return node, true
	}
	if ok2 {
		right = 1
	}

	if currentNode.Val == p.Val || currentNode.Val == q.Val {
		mid = 1
	}

	sum := mid + left + right
	if sum >= 2 {
		return currentNode, true
	} else {
		return nil, sum > 0
	}
}

// 使用map存储一个节点的祖先节点的技巧：节点的左右孩子作为key，节点作为value。
// 这样可以获取到给定节点的父节点，父节点的父节点...也就是给定节点的所有祖先节点。
func helper2(root, p, q *TreeNode) *TreeNode {
	// 帮助遍历二叉树的栈
	stack := make([]*TreeNode, 0)

	// 保存遍历到的节点的父节点
	parents := make(map[*TreeNode]*TreeNode)

	stack = append(stack, root)
	parents[root] = nil

	// 遍历树直到找到p和q
	var node *TreeNode
	ok1, ok2 := containsKey(parents, p), containsKey(parents, q)
	for ; !ok1 || !ok2; ok1, ok2 = containsKey(parents, p), containsKey(parents, q) {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1] // 栈pop
		if node.Left != nil {
			parents[node.Left] = node
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			parents[node.Right] = node
			stack = append(stack, node.Right)
		}
	}

	// p节点的祖先集合
	ancestors := make(map[*TreeNode]bool)

	// 从p开始依次找到所有祖先，相当于构建以p节点为头的链表
	for p != nil {
		ancestors[p] = true
		p = parents[p]
	}

	// 第一个q的祖先且同时是p的祖先的节点，就是p和q的最低公共祖先
	// 这里相当于构建以q为头的链表的同时判断那个节点同时是p链表的节点
	ok := ancestors[q]
	for ; !ok; ok = ancestors[q] {
		q = parents[q]
	}

	return q
}

func containsKey(m map[*TreeNode]*TreeNode, key *TreeNode) bool {
	_, ok := m[key]
	return ok
}
