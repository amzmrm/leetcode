package leetcode

import (
	"container/heap"
)

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	if len(arr) <= k {
		return arr
	}

	partitionArr(arr, 0, len(arr)-1, k)

	return arr[:k]
}

// 划分数组 arr 的 [l,r] 部分，使前 k 小的数在数组的左侧。
func partitionArr(arr []int, left, right, k int) {
	if left >= right {
		return
	}

	pos := partition(arr, left, right)
	pivotIndex := pos - left + 1 // k是从1开始，pivotIndex却是从0开始，所以要+1
	if k == pivotIndex {         // 如果 pos - l + 1 == k，表示 pivot 就是第 k 小的数，直接返回即可；
		return
	} else if k < pivotIndex {
		// 如果 pos - l + 1 < k，表示第 k 小的数在 pivot 的右侧，
		// 因此递归调用 randomized_selected(arr, pos + 1, r, k - (pos - l + 1))；
		partitionArr(arr, left, pos-1, k)
	} else {
		// 如果 pos - l + 1 > k，表示第 k 小的数在 pivot 的左侧，
		// 递归调用 randomized_selected(arr, l, pos - 1, k)。
		partitionArr(arr, pos+1, right, k-pivotIndex)
	}
}

// 随机选一个数pivot，把比pivot小的数放在pivot的左边，比pivot大的数放在右边
func partition(nums []int, left, right int) int {
	i, j, p := left, right, right
	for i != j {
		for nums[i] < nums[p] {
			i++
		}
		for i < j && nums[j] >= nums[p] {
			j--
		}
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[p] = nums[p], nums[i]
	return i
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getLeastNumbers2(arr []int, k int) []int {
	h := IntHeap(arr)
	heap.Init(&h)
	var res []int
	for k > 0 {
		res = append(res, heap.Remove(&h, 0).(int))
		k--
	}
	return res
}
