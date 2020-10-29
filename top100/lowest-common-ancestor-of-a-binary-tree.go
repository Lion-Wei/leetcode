package top100

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var step func(n *TreeNode) int
	var res *TreeNode

	step = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		c := 0
		if n == p {
			c = 1
		}
		if n == q {
			c = 2
		}
		l := step(n.Left)
		r := step(n.Right)
		if l+r+c == 3 {
			res = n
			return 0
		}
		return l + r
	}

	step(root)
	return res
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i > 0; i-- {
		res[i] = res[i+1]
	}
	return nums
}
