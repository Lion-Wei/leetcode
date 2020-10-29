package daily

func sumNumbers(root *TreeNode) int {
	l := []int{}
	var step func(n *TreeNode, num int)

	step = func(n *TreeNode, num int) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			l = append(l, num)
			return
		}
		num = num*10 + n.Val
		step(n.Left, num)
		step(n.Right, num)
	}
	step(root, 0)
	res := 0
	for i := range l {
		res += l[i]
	}
	return res
}
