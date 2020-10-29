package daily

func totalNQueens1(n int) int {
	var step func(x int) int
	var cache []int

	cache = make([]int, n)
	step = func(x int) int {
		if x >= n {
			return 1
		}
		res := 0
		for i := 1; i <= n; i++ {
			canSit := true
			for j := 0; j < x; j++ {
				if cache[j] == i || sub(cache[j], i) == sub(j, x) {
					canSit = false
					break
				}
			}
			if !canSit {
				continue
			}
			cache[x] = i
			res += step(x + 1)
			cache[x] = 0
		}
		return res
	}
	return step(0)
}

func sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - 2
}
