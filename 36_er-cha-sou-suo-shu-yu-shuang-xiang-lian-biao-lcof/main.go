package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Traversal() []int {
	var values []int
	root := t
	for n := t; n != nil; n = n.Right {
		if n.Right == root {
			values = append(values, n.Val)
			break
		}
		values = append(values, n.Val)
	}
	return values
}

func (t *TreeNode) ReverseTraversal() []int {
	var values []int
	root := t
	for n := t; n != nil; n = n.Right {
		if n.Right == root {
			values = append([]int{n.Val}, values...)
			break
		}
		values = append([]int{n.Val}, values...)
	}
	return values
}

func ConvertToDoubleList(node *TreeNode) *TreeNode {
	var lastNodeInList *TreeNode
	lastNodeInList = convertToDoublyList(node, lastNodeInList)

	head := lastNodeInList
	for head != nil && head.Left != nil {
		head = head.Left
	}

	if head != nil {
		head.Left = lastNodeInList
	}
	if lastNodeInList != nil {
		lastNodeInList.Right = head
	}

	return head
}

func convertToDoublyList(node *TreeNode, lastNodeInList *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	current := node
	if current.Left != nil {
		lastNodeInList = convertToDoublyList(current.Left, lastNodeInList)
	}

	current.Left = lastNodeInList
	if lastNodeInList != nil {
		lastNodeInList.Right = current
	}

	lastNodeInList = current

	if current.Right != nil {
		lastNodeInList = convertToDoublyList(current.Right, lastNodeInList)
	}

	return lastNodeInList
}
