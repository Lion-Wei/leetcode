package struct_algorithm

func sortArray(nums []int) []int {
	var sort func(left, right int)

	sort = func(left, right int) {
		if right-left < 1 {
			return
		}
		key, i, j := nums[(left+right)/2], left, right
		for i <= j {
			for ; nums[i] < key; i++ {
			}
			for ; nums[j] > key; j-- {
			}
			if i > j {
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
