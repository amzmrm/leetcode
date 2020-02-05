package problem_191_number_of_1_bits

// 首先把i和1做与运算，判断i的最低位是不是为1。接着把1左移一位得到2，再和i做与运算，就能判断i的次低位是不是1……
// 这样反复左移，每次都能判断i的其中一位是不是1
func HammingWeight(num uint32) int {
	count := 0
	var flag uint32 = 1
	for flag > 0 {
		if num&flag > 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

// 规律：把一个整数减去1，再和原整数做与运算，会把该整数最右边一个1变成0。那么一个整数的二进制表示中有多少个1，就可以进行多少次这样的操作
func HammingWeight2(num uint32) int {
	count := 0
	for num > 0 {
		count++
		num = (num - 1) & num
	}
	return count
}
