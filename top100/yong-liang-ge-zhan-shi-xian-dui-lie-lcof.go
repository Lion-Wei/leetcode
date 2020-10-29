package top100

type CQueue struct {
	head, tail *ListNode
}

func Constructor() CQueue {
	n := &ListNode{}
	return CQueue{
		head: n,
		tail: n,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.tail.Next = &ListNode{
		Val: value,
	}
	this.tail = this.tail.Next
}

func (this *CQueue) DeleteHead() int {
	if this.head.Next == nil {
		return -1
	}
	this.head = this.head.Next
	return this.head.Val
}
