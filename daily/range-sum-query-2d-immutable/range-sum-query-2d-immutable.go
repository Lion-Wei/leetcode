package range_sum_query_2d_immutable

type NumMatrix struct {
	regionSums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	m := NumMatrix{
		regionSums: make([][]int, len(matrix)+1),
	}
	m.regionSums[0] = make([]int, len(matrix[0])+1)
	for i := range matrix {
		m.regionSums[i+1] = make([]int, len(matrix[i])+1)
		for j := range matrix[i] {
			m.regionSums[i+1][j+1] = m.regionSums[i+1][j] + m.regionSums[i][j+1] - m.regionSums[i][j] + matrix[i][j]
		}
	}
	return m
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.regionSums[row2+1][col2+1] + this.regionSums[row1][col1] - this.regionSums[row1][col2+1] - this.regionSums[row2+1][col1]
}
