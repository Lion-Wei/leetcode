package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for right-left > 1 {
		if nums[right] > nums[left] {
			return nums[left]
		}
		mid := (left + right) >> 1
		if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			return nums[mid]
		}
		if nums[mid] < nums[left] {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[left] > nums[right] {
		return nums[right]
	}
	return nums[left]
}
