package leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	var track []int

	backtrack(nums, &res, &track, 0)

	return res
}

func backtrack(nums []int, res *[][]int, track *[]int, start int) {
	newTrack := make([]int, len(*track))
	copy(newTrack, *track)
	*res = append(*res, newTrack)

	for i := start; i < len(nums); i++ {
		*track = append(*track, nums[i])
		backtrack(nums, res, track, i+1)
		*track = (*track)[:len(*track)-1]
	}
}
