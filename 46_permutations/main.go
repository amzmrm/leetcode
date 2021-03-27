package leetcode

func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))
	backtrack(nums, &track, &res, &used)
	return res
}

func backtrack(nums []int, track *[]int, res *[][]int, used *[]bool) {
	if len(*track) == len(nums) { // 遍历到决策树的叶子节点（选择结果的length满足条件），记录得到的选择，结束递归
		t := make([]int, len(*track), cap(*track))
		copy(t, *track)
		*res = append(*res, t)
		track = nil
		return
	}
	for i := 0; i < len(nums); i++ {
		if (*used)[i] { // 使用used记录某个选择是否已经被选择过，避免重复选择
			continue
		}
		*track = append(*track, nums[i]) // 做出选择，选择nums[i]，记录到走过的路径里（track），同时记录i为已经选择过
		(*used)[i] = true

		backtrack(nums, track, res, used) // 做出选择后开始递归，继续在做出选择后的这条路上前进，直到到达叶子节点

		*track = (*track)[:len(*track)-1] // 撤销选择，回到上一步，然后选择其它的节点继续。从而达到能够遍历整个决策树的目的
		(*used)[i] = false
	}
}
