package leetcode

type CQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() CQueue {
	return CQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackIn = append(this.stackIn, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackOut) == 0 {
		for i := len(this.stackIn) - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = nil
	}

	if len(this.stackOut) == 0 {
		return -1
	}

	var head int
	this.stackOut, head = this.stackOut[:len(this.stackOut)-1], this.stackOut[len(this.stackOut)-1]
	return head
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
