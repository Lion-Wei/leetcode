package bytedance

func nextGreaterElements(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	max := 0
	for i := range nums {
		if nums[i] > nums[max] {
			max = i
		}
	}
	s := &stack{s: []int{nums[max]}}
	nums[max] = -1
	for i := max - 1; i >= 0; i-- {
		nums[i] = s.insert(nums[i])
	}
	for i := len(nums) - 1; i > max; i-- {
		nums[i] = s.insert(nums[i])
	}
	return nums
}

type stack struct {
	s []int
}

func (s *stack) insert(n int) int {
	i := len(s.s) - 1
	for ; i >= 0 && s.s[i] <= n; i-- {
	}
	if i < 0 {
		s.s = s.s[:1]
		return -1
	}
	s.s = append(s.s[:i+1], n)
	return s.s[i]
}

func backspaceCompare(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	var m, n int
	for i >= 0 && j >= 0 {
		if S[i] == '#' {
			m++
			i--
			continue
		}
		if T[j] == '#' {
			n++
			j--
			continue
		}
		if m > 0 || n > 0 {
			for m > 0 {
				m--
				i--
			}
			for n > 00 {
				n--
				j--
			}
		} else {
			if S[i] != T[j] {
				return false
			}
			i--
			j--
		}
	}
	return i == j || (i < 0 && j < 0)
}
