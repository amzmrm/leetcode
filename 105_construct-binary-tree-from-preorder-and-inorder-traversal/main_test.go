package leetcode

import "testing"

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	got := buildTree(preorder, inorder)
	if got.Val != 3 {
		t.Fatalf("expect: 3, got: %d", got.Val)
	}
}
