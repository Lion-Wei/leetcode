package top100

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	slice := []int{}
	input := func(n int) {
		i := len(slice) - 1
		for ; i >= 0 && nums[slice[i]] < nums[n]; i-- {
			slice = slice[:i]
		}
		slice = append(slice, n)
	}
	res := make([]int, len(nums)-k+1)
	for i := 0; i < k; i++ {
		input(i)
	}
	list.New()
	res[0] = nums[slice[0]]
	for i := 1; i < len(res); i++ {
		input(i + k - 1)
		if slice[0] == i-1 {
			slice = slice[1:]
		}
		res[i] = nums[slice[0]]
	}
	return res
}
