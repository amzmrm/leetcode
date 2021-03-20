package leetcode

func minArray(numbers []int) int {
	left, right, mid := 0, len(numbers)-1, 0
	for left < right {
		mid = (right + left) / 2
		if numbers[mid] == numbers[right] {
			right--
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		}
	}
	return numbers[left]
}
