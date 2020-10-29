package top100

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if right-left <= 1 {
			break
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] && nums[mid] > nums[right] {
			if target > nums[left] && target < nums[mid] {
				right = mid
			} else {
				left = mid
			}
			continue
		}
		if nums[mid] < nums[left] && nums[mid] < nums[right] {
			if target > nums[mid] && target < nums[right] {
				left = mid
			} else {
				right = mid
			}
			continue
		}
		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	return -1
}
