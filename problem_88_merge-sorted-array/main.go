package main

// 时间复杂度：O((m+n)log(m+n))
// 空间复杂度：O(1)
func Merge0(nums1 []int, m int, nums2 []int, n int) []int {
	for i, num := range nums2 {
		nums1[m+i] = num
	}

	qsort(nums1, 0, len(nums1)-1)

	return nums1
}

// 时间复杂度：O(m+n)
// 空间复杂度：O(m)
func Merge1(nums1 []int, m int, nums2 []int, n int) []int {
	nums1Copy := make([]int, len(nums1))
	copy(nums1Copy, nums1)

	i, j, p := 0, 0, 0
	for i < m && j < n {
		if nums1Copy[i] <= nums2[j] {
			nums1[p] = nums1Copy[i]
			i++
		} else {
			nums1[p] = nums2[j]
			j++
		}
		p++
	}

	if i < m {
		copy(nums1[p:p+m-i], nums1Copy[i:])
	}
	if j < n {
		copy(nums1[p:p+n-j], nums2[j:])
	}

	return nums1
}

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func Merge2(num1 []int, m int, num2 []int, n int) []int {
	i, j := m-1, n-1
	p := len(num1) - 1
	for i >= 0 && j >= 0 {
		if num1[i] <= num2[j] {
			num1[p] = num2[j]
			j--
		} else {
			num1[p] = num1[i]
			i--
		}
		p--
	}
	return num1
}

func qsort(arr []int, left, right int) {
	if left >= right { // 当数组只有一个元素的时候退出递归，只有一个元素qsort(arr, left, p-1)和qsort(arr, p+1, right)会导致数组越界
		return
	}

	p := partition(arr, left, right)

	qsort(arr, left, p-1)  // 递归处理停止位置左边的数列
	qsort(arr, p+1, right) // 递归处理停止位置右边的数列
}

func partition(arr []int, left, right int) int {
	// 以最右边的数字为基准
	i, j, pivot := left, right, arr[right]

	for i != j {
		// 从左边开始，直到找到大于等于基准的数的位置
		for arr[i] < pivot {
			i++
		}

		// 从右边开始，直到找到小于基准的数的位置
		for arr[j] >= pivot && j > i {
			j--
		}

		// 左右标记都停下来了，如果不是在同一个位置停下来，就交换左右位置的数字，从而达到小于基准的数放在左边，大于或等于基准的数放在右边
		if i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 左右标记相遇，本轮对比结束，交换结束位置和基准的数，从而达到基准归位的目的
	// 从下面这行代码就可以看出快排的不稳定：当i=3是，arr[i]==8==pivot，但是pivot和arr[i]却需要对换
	arr[i], arr[right] = pivot, arr[i]

	return i
}
