package daily

func hammingWeight(num uint32) int {
	res := 0
	for ; num != 0; num = num >> 1 {
		if num&1 == 1 {
			res++
		}
	}
	return res
}
