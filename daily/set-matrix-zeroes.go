package daily

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	var firstRow bool
	for i := range matrix[0] {
		if matrix[0][i] == 0 {
			firstRow = true
			break
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 1; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if firstRow {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}
