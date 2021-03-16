package daily

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	value := 1
	for left < right && top < bottom {
		for i := left; i < right; i++ {
			res[top][i] = value
			value++
		}
		for i := top; i < bottom; i++ {
			res[i][right] = value
			value++
		}
		for i := right; i > left; i-- {
			res[bottom][i] = value
			value++
		}
		for i := bottom; i > top; i-- {
			res[i][left] = value
			value++
		}
		left++
		right--
		top++
		bottom--
	}
	if top == bottom {
		for i := left; i <= right; i++ {
			res[top][i] = value
			value++
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			res[i][right] = value
			value++
		}
	}
	return res
}
