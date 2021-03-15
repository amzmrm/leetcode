package leetcode

type MyStack struct {
	queue1 []int
	queue2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	s.queue2 = append(s.queue2, x)
	for i := range s.queue1 {
		s.queue2 = append(s.queue2, s.queue1[i])
	}
	s.queue1 = nil
	s.queue1, s.queue2 = s.queue2, s.queue1
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	top := s.Top()
	s.queue1 = s.queue1[1:]
	return top
}

/** Get the top element. */
func (s *MyStack) Top() int {
	return s.queue1[0]
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
