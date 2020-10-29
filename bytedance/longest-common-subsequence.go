package bytedance

func longestCommonSubsequence(text1 string, text2 string) int {
	leng1, leng2 := len(text1), len(text2)
	dp0 := make([]int, len(text2))
	dp1 := make([]int, len(text2))
	if text2[0] == text1[0] {
		dp0[0] = 1
	}
	for i := 1; i < leng2; i++ {
		if dp0[i-1] == 1 || text1[0] == text2[i] {
			dp0[i] = 1
			continue
		}
	}
	for i := 1; i < leng1; i++ {
		if dp0[0] == 1 || text1[i] == text2[0] {
			dp1[0] = 1
		}
		for j := 1; j < leng2; j++ {
			if text1[i] == text2[j] {
				dp1[j] = dp0[j-1] + 1
			} else {
				dp1[j] = max(dp0[j], dp1[j-1])
			}
		}
		dp0, dp1 = dp1, make([]int, len(text2))
	}
	return dp1[leng2-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
