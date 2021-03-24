package daily

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	leftMin := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= leftMin {
			leftMin = nums[i]
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] && nums[j] > leftMin {
				return true
			}
		}
	}
	return false
}
