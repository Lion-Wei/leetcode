package top100

import "math"

func maxProfit1(prices []int) int {
	leng := len(prices)
	if leng < 2 {
		return 0
	}
	bindMax := make([]int, leng)
	min, x := 0, 0
	for i := 1; i < leng; i++ {
		if prices[i] < prices[min] {
			min = i
		} else if prices[i]-prices[min] > x {
			x = prices[i] - prices[min]
		}
		bindMax[i] = x
	}
	res, min, x := bindMax[leng-1], leng-1, 0
	for i := leng - 2; i >= 0; i-- {
		if prices[i] > prices[min] {
			min = i
		} else if prices[min]-prices[i] > x {
			x = prices[min] - prices[i]
		}
		if x+bindMax[i] > res {
			res = x + bindMax[i]
		}
	}
	return res
}

// 动态规划
func maxProfit(prices []int) int {
	dp10, dp11 := 0, math.MinInt32
	dp20, dp21 := 0, math.MinInt32
	max := func(first, second int) int {
		if first > second {
			return first
		}
		return second
	}
	// base case:
	// dp[-1][k][0] = dp[i][0][0] = 0
	// dp[-1][k][1] = dp[i][0][1] = -infinity
	// 状态转移方程：
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
	for i := range prices {
		// only count buying as a transation
		// use multiple assignment to avoid sequence problem
		dp20, dp21, dp10, dp11 = max(dp20, dp21+prices[i]), max(dp21, dp10-prices[i]), max(dp10, dp11+prices[i]), max(dp11, -prices[i])
	}
	return dp20
}
