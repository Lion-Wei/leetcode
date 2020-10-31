package top100

func moveZeroes(nums []int) {
	i, count := 0, 0
	for i < len(nums) {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			count++
		} else {
			i++
		}
	}
	nums = append(nums, make([]int, count)...)
}
