package daily

func deleteDuplicates(head *ListNode) *ListNode {
	dumpy := &ListNode{Next: head}
	head = dumpy
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal := head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dumpy.Next
}
