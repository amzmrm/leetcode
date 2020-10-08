package main

import "testing"

func TestGetLastCommonAncestor(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		root := TreeNode{Val: "A"}
		nodeB := TreeNode{Val: "B"}
		nodeC := TreeNode{Val: "C"}
		nodeD := TreeNode{Val: "D"}
		nodeF := TreeNode{Val: "F"}
		nodeG := TreeNode{Val: "G"}
		nodeE := TreeNode{Val: "E"}
		nodeH := TreeNode{Val: "H"}
		nodeI := TreeNode{Val: "I"}
		nodeJ := TreeNode{Val: "J"}

		root.Children = append(root.Children, &nodeB, &nodeC)
		nodeB.Children = append(nodeB.Children, &nodeD, &nodeE)
		nodeD.Children = append(nodeD.Children, &nodeF, &nodeG)
		nodeE.Children = append(nodeE.Children, &nodeH, &nodeI, &nodeJ)

		got := GetLastCommonAncestor(&root, &nodeF, &nodeH)
		if got == nil {
			t.Fatalf("got:%v, expect:%v", nil, "B")
		}
		if got.Val != "B" {
			t.Fatalf("got:%v, expect:%v", got.Val, "B")
		}
	})

	t.Run("nodes are direct children of root node", func(t *testing.T) {
		root := TreeNode{Val: "A"}
		nodeB := TreeNode{Val: "B"}
		nodeC := TreeNode{Val: "C"}
		nodeD := TreeNode{Val: "D"}
		nodeF := TreeNode{Val: "F"}
		nodeG := TreeNode{Val: "G"}
		nodeE := TreeNode{Val: "E"}
		nodeH := TreeNode{Val: "H"}
		nodeI := TreeNode{Val: "I"}
		nodeJ := TreeNode{Val: "J"}

		root.Children = append(root.Children, &nodeB, &nodeC)
		nodeB.Children = append(nodeB.Children, &nodeD, &nodeE)
		nodeD.Children = append(nodeD.Children, &nodeF, &nodeG)
		nodeE.Children = append(nodeE.Children, &nodeH, &nodeI, &nodeJ)

		got := GetLastCommonAncestor(&root, &nodeB, &nodeC)
		if got == nil {
			t.Fatalf("got:%v, expect:%v", nil, "A")
		}
		if got.Val != "A" {
			t.Fatalf("got:%v, expect:%v", got.Val, "A")
		}
	})

	t.Run("case2", func(t *testing.T) {
		root := TreeNode{Val: "A"}
		nodeB := TreeNode{Val: "B"}
		nodeC := TreeNode{Val: "C"}
		nodeD := TreeNode{Val: "D"}
		nodeF := TreeNode{Val: "F"}
		nodeG := TreeNode{Val: "G"}
		nodeE := TreeNode{Val: "E"}
		nodeH := TreeNode{Val: "H"}
		nodeI := TreeNode{Val: "I"}
		nodeJ := TreeNode{Val: "J"}

		root.Children = append(root.Children, &nodeB, &nodeC)
		nodeB.Children = append(nodeB.Children, &nodeD, &nodeE)
		nodeD.Children = append(nodeD.Children, &nodeF, &nodeG)
		nodeE.Children = append(nodeE.Children, &nodeH, &nodeI, &nodeJ)

		got := GetLastCommonAncestor(&root, &root, &nodeJ)
		if got == nil {
			t.Fatalf("got:%v, expect:%v", nil, "A")
		}
		if got.Val != "A" {
			t.Fatalf("got:%v, expect:%v", got.Val, "A")
		}
	})
}
