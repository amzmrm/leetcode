package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) { // 确保第一个数组的len是两者中小的那个
		return findMedianSortedArrays(nums2, nums1)
	}

	// 思路
	// 使用一条分割线把两个数组分别分割成两部分，使分割线左边的部分和右边的部分同时满足以下条件：
	// 1. 分割线左边和右边的元素个数相等（偶数），或者左边元素个数比右边元素个数多1个（奇数）；
	// 2. 分割线左边的所有元素的数值 <= 分割线右边的所有元素的数组。
	// 当m+n为偶数的时候，分割线左边元素个数是(m+n)/2，也可以写出(m+n+1)/2；
	// 当m+n为奇数的时候，分割线右边元素个数是(m+n+1)/2。为了编码方便统一写成(m+n+1)/2

	// 找到这条返回条件的分割线后，这条分割线的性质就可以和只有一个有序数组的情况同意起来：中位数只和分割线两侧的4个元素有关，
	// 当两个有序数组的长度之和为奇数的时候，分割线左边两个数值中的最大值就是数组的中位数；
	// 当两个有序数组的长度之和为偶数的时候，分割线左边两个数值中的最大值就是其中一个中位数，分割线右边两个数值中的最小值就是另一个中位数

	m, n := len(nums1), len(nums2)
	totalLeft := (m + n + 1) / 2 // 分割线左边的所有元素需要满足的个数

	// 在nums1的区间[0,m]里查找符合条件的分割线
	// 使得nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i]
	left, right := 0, m
	for left < right {
		i := (right + left + 1) / 2
		j := totalLeft - i
		if nums1[i-1] > nums2[j] {
			// 分割线在nums1上太靠右了，需要往左缩小，下一轮搜索区间[left, i-1]
			right = i - 1
		} else {
			// 下一轮搜索区间[i, right]
			left = i
		}
	}

	median1 := left                // 分割线在第一个数组的下标
	median2 := totalLeft - median1 // 分割线在第二个数组的下标

	nums1LeftMax := math.MinInt32
	if median1 != 0 {
		nums1LeftMax = nums1[median1-1]
	}
	nums1RightMin := math.MaxInt32
	if median1 != m {
		nums1RightMin = nums1[median1]
	}

	nums2LeftMax := math.MinInt32
	if median2 != 0 {
		nums2LeftMax = nums2[median2-1]
	}
	nums2RightMin := math.MaxInt32
	if median2 != n {
		nums2RightMin = nums2[median2]
	}

	if (m+n)%2 == 1 {
		// 当两个有序数组的长度之和为奇数的时候，分割线左边两个数值中的最大值就是数组的中位数
		return float64(max(nums1LeftMax, nums2LeftMax))
	} else {
		// 当两个有序数组的长度之和为偶数的时候，分割线左边两个数值中的最大值就是其中一个中位数，分割线右边两个数值中的最小值就是另一个中位数
		return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 解法2，时间复制度O(log(m+n))
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		midIndex := totalLen / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLen/2-1, totalLen/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k = k - (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k = k - (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}
