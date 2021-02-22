package daily

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return true
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]);j++{
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
