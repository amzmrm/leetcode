package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	var first, last *Node
	var fn func(node *Node)
	fn = func(node *Node) {
		if node == nil {
			return
		}
		fn(node.Left)
		if last != nil {
			last.Right = node
			node.Left = last
		} else {
			first = node
		}
		last = node
		fn(node.Right)
	}

	fn(root)
	last.Right = first
	first.Left = last
	return first
}
