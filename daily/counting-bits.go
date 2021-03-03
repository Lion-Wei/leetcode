package daily

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	base := 1
	bits := make([]int, num+1)
	for i := 1; i < len(bits); i++ {
		if i-base == base {
			bits[i] = 1
			base = i
		} else {
			bits[i] = 1 + bits[i-base]
		}
	}
	return bits
}

func countBits1(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}
