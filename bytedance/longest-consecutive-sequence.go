package bytedance

func longestConsecutive(nums []int) int {
	s := make(map[int]bool)
	for i := range nums {
		s[nums[i]] = true
	}
	res := 0
	for i := range s {
		if s[i-1] {
			continue
		}
		n, long := i, 0
		for s[n] {
			n++
			long++
		}
		if long > res {
			res = long
		}
	}
	return res
}
