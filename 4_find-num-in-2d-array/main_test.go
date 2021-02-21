package leetcode

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {
	type test struct {
		name   string
		matrix [][]int
		target int
		expect bool
	}

	tests := []test{
		{
			name: "normal case",
			matrix: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			target: 7,
			expect: true,
		},
		{
			name: "another case",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 8, 16, 22},
				{10, 13, 14, 17, 4},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			expect: true,
		},
		{
			name: "another case 2",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 8, 16, 22},
				{10, 13, 14, 17, 4},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			expect: false,
		},
		{
			name: "length == 1",
			matrix: [][]int{
				{4, 5, 6, 7, 8, 9},
			},
			target: 4,
			expect: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FindNumberIn2DArray(test.matrix, test.target)
			if got != test.expect {
				t.Fatalf("got:%v, expect:%v", got, test.expect)
			}
		})
	}
}
