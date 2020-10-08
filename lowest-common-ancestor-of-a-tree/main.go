package main

type TreeNode struct {
	Val      string
	Children []*TreeNode
}

// 得到从根结点pRoot开始到达结点pNode的路径，这条路径保存在path中
// 递归公式：
// add currRoot to path;
// if (currRoot.Child1 contains node) { add node to path } else { remove currRoot from path}
// 递归结束的条件：currRoot.Val == node.Val
func GetNodePath(root *TreeNode, node *TreeNode) []*TreeNode {
	if root.Val == node.Val {
		// 当前的root是目标节点，返回它
		return []*TreeNode{root}
	}

	// 先把当前的root节点加到path中
	path := append([]*TreeNode{}, root)

	var chilePath []*TreeNode
	for i := 0; i < len(root.Children) && len(chilePath) == 0; i++ {
		chilePath = GetNodePath(root.Children[i], node)
	}

	if len(chilePath) == 0 {
		// 目标节点不在当前root节点的子树，把先前加进去的当前root节点从path中移除
		path = path[:len(path)-1]
	} else {
		// 目标节点在当前root节点的子树，把当前root节点加到path中
		path = append(path, chilePath...)
	}

	return path
}

// 得到两个路径path1和path2的最后一个公共结点
func GetLastCommonNode(path1 []*TreeNode, path2 []*TreeNode) *TreeNode {
	var commonNode *TreeNode

	for i, j := 0, 0; i < len(path1) && j < len(path2); {
		if path1[i].Val == path2[j].Val {
			commonNode = path1[i]
		}
		i++
		j++
	}

	return commonNode
}

// 先调用GetNodePath得到pRoot到达pNode1的路径path1，再得到pRoot到达pNode2的路径path2，
// 接着调用GetLastCommonPath得到path1和path2的最后一个公共结点，即我们要找的最低公共祖先
func GetLastCommonAncestor(root *TreeNode, node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if root == nil || node1 == nil || node2 == nil {
		return nil
	}

	path1 := GetNodePath(root, node1)
	path2 := GetNodePath(root, node2)

	return GetLastCommonNode(path1, path2)
}
