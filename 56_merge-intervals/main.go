package leetcode

import "sort"

func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 先将区间根据起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > merged[len(merged)-1][1] {
			// 区间的开始坐标大于最后一个已合并区间的结束坐标，表示这两个区间没有重叠
			merged = append(merged, intervals[i])
			continue
		} else if intervals[i][0] <= merged[len(merged)-1][1] {
			// 区间的开始坐标小于或等于最后一个已合并区间的结束坐标，表示这两个区间有重叠
			if intervals[i][1] > merged[len(merged)-1][1] {
				// 区间的结束坐标大于最后一个已合并区间的结束坐标，表示要合并
				merged[len(merged)-1][1] = intervals[i][1]
			}
		}
	}
	return merged
}
