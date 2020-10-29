package top100

func numSquares(n int) int {
	squares := make([]int, n+1)
	squares[0] = 0
	for i := 1; i <= n; i++ {
		squares[i] = i
		for j := 1; ; j++ {
			if i-j*j < 0 {
				break
			}
			squares[i] = min(squares[i], squares[i-j*j]+1)
		}
	}
	return squares[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
