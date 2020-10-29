package top100

func reverseKGroup1(head *ListNode, k int) *ListNode {
	if head.Next == nil || k < 2 {
		return head
	}
	pre := &ListNode{
		Next: head,
	}
	kHead := pre
	for n := kHead.Next; n.Next != nil; kHead, n = n, n.Next {
		i := 1
		for ; i < k && n.Next != nil; i++ {
			tmp := n.Next
			n.Next, tmp.Next, kHead.Next = tmp.Next, kHead.Next, tmp
		}
		if i < k {
			for n = kHead.Next; n.Next != nil; {
				tmp := n.Next
				n.Next, tmp.Next, kHead.Next = tmp.Next, kHead.Next, tmp
			}
			break
		}
	}
	return pre.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	//查看是否有k个节点
	var p = head
	for i := 0; i < k; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	//翻转前k个节点
	var p1, p2 = head, head.Next
	for i := 0; i < k-1; i++ {
		var nextP2 = p2.Next
		p2.Next, p1, p2 = p1, p2, nextP2
	}
	//递归翻转剩余链表
	head.Next = reverseKGroup(p2, k)
	return p1
}
