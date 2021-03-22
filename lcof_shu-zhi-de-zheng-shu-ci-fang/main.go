package leetcode

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n < 0 {
		return 1 / helper(x, -n)
	}
	return helper(x, n)
}

func helper(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		half := myPow(x, n/2)
		return half * half
	} else {
		half := myPow(x, n/2)
		return half * half * x
	}
}
