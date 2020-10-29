package top100

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/submissions/

/*
滑动窗口，时间复杂度最坏情况下可达到n^2/2
可以考虑优先级队列的方式进行优化
*/
// TODO： 优先级队列
type sortList struct {
	compare func(a, b int) bool
	buffer  *ListNode
	size    int
}

func newSortList(size int, compare func(a, b int) bool) *sortList {
	return &sortList{
		compare: compare,
		buffer:  &ListNode{},
		size:    size,
	}
}

func (s *sortList) insert(n int) {
	tmp := s.buffer
	for ; tmp.Next != nil && s.compare(n, tmp.Next.Val); tmp = tmp.Next {
	}
	tmp.Next = &ListNode{
		Val:  n,
		Next: tmp.Next,
	}
	if s.size > 0 {
		s.size--
	} else {
		s.buffer = s.buffer.Next
	}
}

func findKthLargest1(nums []int, k int) int {
	compare := func(a, b int) bool {
		return a > b
	}
	leng := len(nums)
	if k > leng-k+1 {
		k = leng - k + 1
		compare = func(a, b int) bool {
			return a < b
		}
	}

	sl := newSortList(k, compare)

	for _, n := range nums {
		sl.insert(n)
	}
	return sl.buffer.Next.Val
}

/*
基于快速排序
*/
func findKthLargest(nums []int, k int) int {
	var sort func(left, right int)

	sort = func(l, r int) {
		if r < l+1 {
			return
		}
		pivot := (l + r) >> 1
		nums[pivot], nums[r] = nums[r], nums[pivot]
		i, j, m := l-1, l, nums[r]
		for ; j <= r; j++ {
			if nums[j] <= m {
				i++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		switch {
		case k == i:
			return
		case k < i:
			sort(l, i-1)
		case k > i:
			sort(i+1, r)
		}
	}
	k = len(nums) - k
	sort(0, len(nums)-1)
	return nums[k]
}
