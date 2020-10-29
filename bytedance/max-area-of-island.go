package bytedance

func maxAreaOfIsland(grid [][]int) int {
	var step func(i, j int) int

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	leng1, leng2 := len(grid), len(grid[0])
	step = func(i, j int) int {
		if i < 0 || j < 0 || i >= leng1 || j >= leng2 || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + step(i-1, j) + step(i, j+1) + step(i+1, j) + step(i, j-1)
	}

	var res int
	for i := range grid {
		for j := range grid[i] {
			res = max(step(i, j), res)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
