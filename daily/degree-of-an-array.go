package daily

// daily 2021-02-20
func findShortestSubArray(nums []int) int {
	type el struct {
		start int
		end int
		num int
	}
	res := &el{}
	numCache := make(map[int]*el)

	for i, n := range nums {
		_, found := numCache[n]
		if !found {
			numCache[n] = &el{
				start: i,
			}
		}
		numCache[n].end = i+1
		numCache[n].num = numCache[n].num+1
		switch {
		case numCache[n].num > res.num:
			res = numCache[n]
		case numCache[n].num == res.num && numCache[n].end - numCache[n].start < res.end - res.start:
			res = numCache[n]
		}
	}
	return res.end-res.start
}
