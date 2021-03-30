package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for n := head; n != nil && n.Next != nil; n = n.Next {
		for ; n.Next != nil && n.Next.Val == n.Val; n.Next = n.Next.Next {
		}
	}
	return head
}
