package problem_509_fibonacci_number

func Fib(N int) int {
	results := []int{0, 1}
	if N < 2 {
		return results[N]
	}

	fibOne, fibTwo, fibN := 1, 0, 0
	for i := 2; i <= N; i++ {
		fibN = fibOne + fibTwo
		fibTwo = fibOne
		fibOne = fibN
	}
	return fibN
}
