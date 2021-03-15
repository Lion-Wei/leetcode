package daily

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	res := []int{}

	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top > bottom && left < right {
		res = append(res, matrix[top][left:right]...)
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- {
			res = append(res, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			res = append(res, matrix[i][left])
		}
		top++
		bottom--
		left++
		right--
	}
	if top == bottom {
		res = append(res, matrix[top][left:right+1]...)
	} else if left == right {
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
	}
	return res
}
