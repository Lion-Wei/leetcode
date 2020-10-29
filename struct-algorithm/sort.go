package struct_algorithm

func fastSort(nums []int) []int {
	var sort func(start, end int)

	sort = func(left, right int) {
		if left >= right {
			return
		}
		tmp := nums[(left+right)/2]
		i, j := left, right
		for {
			for ; nums[i] < tmp; i++ {
			}
			for ; nums[j] > tmp; j-- {
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		sort(left, j)
		sort(i, right)
	}
	sort(0, len(nums)-1)
	return nums
}
