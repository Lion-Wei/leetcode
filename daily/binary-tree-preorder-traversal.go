package daily

// 递归
func preorderTraversal1(root *TreeNode) []int {
	var step func(n *TreeNode)

	res := make([]int, 0)

	step = func(n *TreeNode) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		step(n.Left)
		step(n.Right)
	}

	step(root)
	return res
}

// 迭代
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	n, stack := root, []*TreeNode{}
	for n != nil || len(stack) != 0 {
		for n != nil {
			res = append(res, n.Val)
			stack = append(stack, n)
			n = n.Left
		}
		n = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
