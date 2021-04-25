package daily

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	dpf := func(target int) int {
		res := 0
		for i := range nums {
			if target-nums[i] >= 0 {
				res += dp[target-nums[i]]
			}
		}
		return res
	}
	for i := 1; i <= target; i++ {
		dp[i] = dpf(i)
	}
	return dp[target]
}
