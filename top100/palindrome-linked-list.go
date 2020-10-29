package top100

func isPalindrome(head *ListNode) bool {
	n1, n2 := head, head
	var tmp *ListNode
	for n2 != nil && n2.Next != nil {
		n1.Next, tmp, n1 = tmp, n1, n1.Next
		n2 = n2.Next.Next
	}
	if n2 == nil {
		n1, n2 = tmp, n1
	} else {
		n1, n2 = tmp, n1.Next
	}
	for ; n1 != nil && n2 != nil; n1, n2 = n1.Next, n2.Next {
		if n1.Val != n2.Val {
			return false
		}
	}
	return true
}
