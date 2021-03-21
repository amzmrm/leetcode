package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	var res int
	var nums []int
	var left, right int
	var num int
	for _, token := range tokens {
		if isOp(token) {
			nums, left, right = nums[:len(nums)-2], nums[len(nums)-2], nums[len(nums)-1]
			res = do(token, left, right)
			nums = append(nums, res)
		} else {
			num = toInt(token)
			res = num
			nums = append(nums, num)
		}
	}
	return res
}

func isOp(op string) bool {
	switch op {
	case "+", "-", "*", "/":
		return true
	}
	return false
}

func do(op string, left, right int) int {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	return -1
}

func toInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
