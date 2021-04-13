package daily

import "math"

func minDiffInBST(root *TreeNode) int {
	pre := math.MinInt32
	res := math.MaxInt32
	var step func(root *TreeNode)
	step = func(root *TreeNode) {
		if root == nil {
			return
		}
		step(root.Left)
		if root.Val-pre < res {
			res = root.Val - pre
		}
		pre = root.Val
		step(root.Right)
	}
	step(root)
	return res
}
