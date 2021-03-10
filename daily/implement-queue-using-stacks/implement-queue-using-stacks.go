package implement_queue_using_stacks

type MyQueue struct {
	l []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		l: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.l = append(this.l, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	re := this.l[0]
	this.l = this.l[1:]
	return re
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.l[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.l) == 0
}
