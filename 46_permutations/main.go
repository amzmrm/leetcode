package leetcode

func permute(nums []int) [][]int {
	var res [][]int

	var track []int
	var backtrack func([]int, []int)
	used := make([]bool, len(nums))
	backtrack = func(ns []int, track []int) {
		if len(track) == len(nums) {
			newTrack := make([]int, len(track), cap(track))
			copy(newTrack, track)
			res = append(res, newTrack)
			return
		}

		for i := 0; i < len(ns); i++ {
			if used[i] {
				continue
			}

			track = append(track, ns[i])
			used[i] = true
			backtrack(ns, track)
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	backtrack(nums, track)

	return res
}
