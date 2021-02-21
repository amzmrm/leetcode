package leetcode

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if cpm, ok := m[complement]; ok {
			return []int{cpm, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
