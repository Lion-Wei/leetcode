package top100

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	points, res := [2][]int{}, 0
	points[0] = make([]int, len(matrix[0])+1)
	points[1] = make([]int, len(matrix[0])+1)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				continue
			}
			n := j + 1
			points[1][n] = min(points[0][j]+1, points[0][n]+1, points[1][j]+1)
			if points[1][n] > res {
				res = points[1][n]
			}
		}
		points[0] = points[1]
		points[1] = make([]int, len(matrix[0])+1)
	}
	return res * res
}

func min(a, b, c int) int {
	res := a
	if res > b {
		res = b
	}
	if res > c {
		res = c
	}
	return res
}
