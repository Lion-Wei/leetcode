package top100

func rightSideView(root *TreeNode) []int {
	var step func(n *TreeNode, depth int)

	res := make([]int, 0)

	step = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}
		if len(res) <= depth {
			res = append(res, n.Val)
		}
		step(n.Right, depth+1)
		step(n.Left, depth+1)
	}
	step(root, 0)
	return res
}
