package clumsy_factorial

func clumsy(N int) int {
	res, tmp := 0, N
	for i := 1; i < N; i++ {
		switch i % 4 {
		case 1:
			tmp = tmp * (N - i)
		case 2:
			tmp = tmp / (N - i)
		case 3:
			res = res + tmp + N - i
		case 0:
			tmp = i - N
		}
	}
	if N%4 != 0 {
		return res + tmp
	}
	return res
}
