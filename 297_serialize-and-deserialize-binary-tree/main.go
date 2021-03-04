package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	Separator string
	Null      string
}

func Constructor() Codec {
	return Codec{
		Separator: ",",
		Null:      "#",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.Null + this.Separator
	}

	var res string
	res += fmt.Sprintf("%d", root.Val) + this.Separator

	left := this.serialize(root.Left)
	res += left
	right := this.serialize(root.Right)
	res += right
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, this.Separator)
	var nodes []string
	for _, str := range vals {
		if len(str) == 0 {
			continue
		}
		nodes = append(nodes, str)
	}

	var first string
	var fn func() *TreeNode
	fn = func() *TreeNode {
		if len(nodes) == 0 {
			return nil
		}

		first, nodes = nodes[0], nodes[1:]
		if first == this.Null {
			return nil
		}
		root := TreeNode{Val: str2Int(first)}
		root.Left = fn()
		root.Right = fn()
		return &root
	}

	return fn()
}

func str2Int(in string) int {
	val, _ := strconv.Atoi(in)
	return val
}
