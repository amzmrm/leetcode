package leetcode

func Fib(N int) int {
	results := []int{0, 1}
	if N < 2 {
		return results[N]
	}

	fibOne, fibTwo := 0, 1
	var fibN int
	for i := 2; i <= N; i++ {
		fibN = fibOne + fibTwo
		fibOne = fibTwo
		fibTwo = fibN
	}
	return fibN
}
