package leetcode

type MinStack struct {
	stack     []int
	minRecord []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:     make([]int, 0),
		minRecord: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, x)
		this.minRecord = append(this.minRecord, x)
		return
	}

	this.stack = append(this.stack, x)
	if len(this.minRecord) == 0 || x <= this.minRecord[len(this.minRecord)-1] {
		this.minRecord = append(this.minRecord, x)
	}
}

func (this *MinStack) Pop() int {
	var top int
	this.stack, top = this.stack[:len(this.stack)-1], this.stack[len(this.stack)-1]
	if top == this.minRecord[len(this.minRecord)-1] {
		this.minRecord = this.minRecord[:len(this.minRecord)-1]
	}
	return top
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minRecord[len(this.minRecord)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
