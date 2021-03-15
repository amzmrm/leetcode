package leetcode

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
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

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stackOut) == 0 {
		for i := len(this.stackIn) - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = nil
	}
	if len(this.stackOut) == 0 {
		return -1
	}
	return this.stackOut[len(this.stackOut)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stackOut) == 0 && len(this.stackIn) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
