package remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i, count := 1, 1
	for i < len(nums)-1 {
		if nums[i] == nums[i-1] {
			if count == 2 {
				nums = append(nums[:i], nums[i+1:]...)
				continue
			}
			count++
		} else {
			count = 1
		}
		i++
	}
	if count == 2 && nums[i] == nums[i-1] {
		nums = nums[:i]
	}
	return len(nums)
}
