package top100

type ListNode struct {
	Val  int
	Next *ListNode
}

type DualListNode struct {
	Val  int
	Next *DualListNode
	Pre  *DualListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildList(nums []int) *ListNode {
	h := &ListNode{}
	n := h
	for i := range nums {
		n.Next = &ListNode{Val: nums[i]}
		n = n.Next
	}
	return h.Next
}

func convertList(l *ListNode) []int {
	res := []int{}
	for ; l != nil; l = l.Next {
		res = append(res, l.Val)
	}
	return res
}
