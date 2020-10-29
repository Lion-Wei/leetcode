package top100

func sortArray(nums []int) []int {
	var sort func(l, r int)

	sort = func(l, r int) {
		if r-l <= 1 {
			return
		}
		pivot := (l + r) >> 1
		nums[pivot], nums[r] = nums[r], nums[pivot]
		i, j, m := l-1, l, nums[r]
		for ; j <= r; j++ {
			if nums[j] <= m {
				i++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		sort(l, i-1)
		sort(i+1, r)
	}

	sort(0, len(nums)-1)
	return nums
}
