package top100

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(n *TreeNode) {
	if n == nil {
		return
	}
	invert(n.Left)
	invert(n.Right)
	n.Left, n.Right = n.Right, n.Left
}
