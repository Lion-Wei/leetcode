package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}
	n := newHead
	for i := 0; i < left-1 && n.Next != nil; n, i = n.Next, i+1 {
	}
	leftNode := n.Next
	for i := 0; i < (right - left); i++ {
		tmp := leftNode.Next
		leftNode.Next = tmp.Next
		tmp.Next = n.Next
		n.Next = tmp
	}
	return newHead.Next
}
