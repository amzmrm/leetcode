package leetcode

func tribonacci(n int) int {
	result := []int{0, 1, 1}
	if n < len(result) {
		return result[n]
	}

	fib1, fib2, fib3 := 0, 1, 1
	var fibN int
	for i := 3; i <= n; i++ {
		fibN = fib1 + fib2 + fib3
		fib1 = fib2
		fib2 = fib3
		fib3 = fibN
	}
	return fibN
}
