package bytedance

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{
		Next: l1,
	}
	x, n1, n2 := 0, res, l2
	for ; n1.Next != nil && n2 != nil; n1, n2 = n1.Next, n2.Next {
		x, n1.Next.Val = sub(n1.Next.Val, n2.Val, x)
	}
	if n1.Next == nil {
		n1.Next = n2
	}
	for ; n1.Next != nil && x != 0; n1 = n1.Next {
		x, n1.Next.Val = sub(n1.Next.Val, x, 0)
	}
	if x != 0 {
		n1.Next = &ListNode{
			Val: x,
		}
	}
	return res.Next
}

func sub(a, b, c int) (int, int) {
	return (a + b + c) / 10, (a + b + c) % 10
}
