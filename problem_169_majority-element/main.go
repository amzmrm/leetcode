package problem_169_majority_element

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func majorityElement(nums []int) int {
	counts := make(map[int]int, 0)
	var count int
	var ok bool
	for _, num := range nums {
		if count, ok = counts[num]; ok {
			count++
			counts[num] = count
		} else {
			counts[num] = 1
		}
	}

	var ret int
	for val := range counts {
		if counts[val] > ret {
			ret = counts[val]
			count = val
		}
	}

	return count
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func majorityElement2(nums []int) int {
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count += -1
		}
	}
	return candidate
}
