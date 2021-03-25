package leetcode

type TwoSum struct {
	numbers map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{numbers: make(map[int]int)}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.numbers[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for number, count := range this.numbers {
		complement := value - number
		if number != complement {
			if _, ok := this.numbers[complement]; ok {
				return true
			}
		} else {
			if count > 1 {
				return true
			}
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
