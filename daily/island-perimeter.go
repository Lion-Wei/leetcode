package daily

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res, width, height := 0, len(grid[0]), len(grid)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			if i == 0 || grid[i-1][j] == 0 {
				res++
			}
			if j == 0 || grid[i][j-1] == 0 {
				res++
			}
			if i+1 == height || grid[i+1][j] == 0 {
				res++
			}
			if j+1 == width || grid[i][j+1] == 0 {
				res++
			}
		}
	}
	return res
}
