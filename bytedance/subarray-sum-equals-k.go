package bytedance

func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	var sum, res int
	m[0] = 1
	for i := range nums {
		sum += nums[i]
		if _, ok := m[sum-k]; ok {
			res += m[sum-k]
		}
		m[sum] += 1
	}
	return res
}

func addStrings(num1 string, num2 string) string {
	res := ""
	var x uint8
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var a, b uint8
		if i >= 0 {
			a = num1[i] - '0'
		}
		if j >= 0 {
			b = num2[j] - '0'
		}
		sum := a + b + x
		x = sum / 10
		res = string(sum%10) + res
	}
	if x > 0 {
		res = string(x+'0') + res
	}
	return res
}
