package daily

func isMonotonic(A []int) bool {
	i := 1
	for ; i < len(A) && A[i] == A[i-1]; i++ {
	}
	if i == len(A) {
		return true
	}
	inc := A[i] > A[i-1]
	for i = i + 1; i < len(A); i++ {
		if inc && A[i] < A[i-1] {
			return false
		}
		if !inc && A[i] > A[i-1] {
			return false
		}
	}
	return true
}
