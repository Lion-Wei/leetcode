package volume_of_histogram_lcci

func trap(height []int) int {
	length := len(height)
	if length < 3 {
		return 0
	}

	left := make([]int, length)
	for i := 1; i < length; i++ {
		left[i] = max(left[i-1], height[i-1])
	}
	var res, right int
	for i := length - 2; i > 0; i-- {
		right = max(right, height[i+1])
		top := min(left[i], right)
		if top < height[i] {
			continue
		}
		res = res + (top - height[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
