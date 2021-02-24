package daily

func flipAndInvertImage(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}
	leng := len(A[0])
	odd := leng%2 == 1
	for i := range A {
		for j := 0; j < leng>>1; j++ {
			A[i][j], A[i][leng-1-j] = invert(A[i][leng-1-j]), invert(A[i][j])
		}
		if odd {
			A[i][leng>>1] = invert(A[i][leng>>1])
		}
	}
	return A
}

func invert(n int) int {
	return n ^ 1
}
