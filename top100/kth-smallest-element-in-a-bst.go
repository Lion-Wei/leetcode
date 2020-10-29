package top100

func kthSmallest(root *TreeNode, k int) int {
	var step func(n *TreeNode)
	var index, res int

	step = func(n *TreeNode) {
		if n == nil {
			return
		}
		step(n.Left)
		index++
		if index == k {
			res = n.Val
			return
		}
		step(n.Right)
	}

	step(root)
	return res
}
