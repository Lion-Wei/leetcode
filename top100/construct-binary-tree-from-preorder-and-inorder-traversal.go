package top100

func buildTree(preorder []int, inorder []int) *TreeNode {
	leng := len(preorder)
	if leng == 0 {
		return nil
	}
	i := 0
	for ; inorder[i] != preorder[0]; i++ {
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:i+1], inorder[:i]),
		Right: buildTree(preorder[i+1:], inorder[i+1:]),
	}
}
