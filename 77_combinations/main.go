package leetcode

func combine(n int, k int) [][]int {
	var track []int
	var res [][]int

	backtrack(n, k, &res, &track, 1)

	return res
}

func backtrack(n int, k int, res *[][]int, track *[]int, start int) {
	if len(*track) == k {
		t := make([]int, k)
		copy(t, *track)
		*res = append(*res, t)
		return
	}

	for i := start; i <= n; i++ {
		*track = append(*track, i)
		backtrack(n, k, res, track, i+1)
		*track = (*track)[:len(*track)-1]
	}
}
