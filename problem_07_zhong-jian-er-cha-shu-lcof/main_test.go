package problem_07_zhong_jian_er_cha_shu_lcof

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type test struct {
		name     string
		preorder []int
		inorder  []int
		expect   []int
	}

	tests := []test{
		{
			name:     "normal case",
			preorder: []int{1, 2, 4, 7, 3, 5, 6, 8},
			inorder:  []int{4, 7, 2, 1, 5, 3, 8, 6},
			expect:   []int{1, 2, 4, 7, 3, 5, 6, 8},
		},
		{
			name:     "another case",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expect:   []int{3, 9, 20, 15, 7},
		},
		{
			name:     "single node tree",
			preorder: []int{1},
			inorder:  []int{1},
			expect:   []int{1},
		},
		{
			name:     "all nodes are right child",
			preorder: []int{1, 2, 4, 7, 3, 5, 6, 8},
			inorder:  []int{1, 2, 4, 7, 3, 5, 6, 8},
			expect:   []int{1, 2, 4, 7, 3, 5, 6, 8},
		},
		{
			name:     "all nodes are left child",
			preorder: []int{1, 2, 4, 7, 3, 5, 6, 8},
			inorder:  []int{8, 6, 5, 3, 7, 4, 2, 1},
			expect:   []int{1, 2, 4, 7, 3, 5, 6, 8},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := BuildTree(test.preorder, test.inorder)
			if !reflect.DeepEqual(PreorderTraversal(got), test.expect) {
				t.Fatalf("got: %v, expect: %v", PreorderTraversal(got), test.expect)
			}
		})
	}
}
