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
