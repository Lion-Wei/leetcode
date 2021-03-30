package daily

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	left, right := 0, len(matrix)*len(matrix[0])-1
	col := len(matrix[0])
	for left != right {
		mid := (left + right) >> 1
		n := matrix[mid/col][mid%col]
		switch {
		case n == target:
			return true
		case n > target:
			right = mid
		case n < target:
			left = mid + 1
		}
	}
	return matrix[left/col][right%col] == target
}
