package top100

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_fastSort(t *testing.T) {
	nums := []int{}
	for i := 0; i < 2; i++ {
		nums = append(nums, rand.Intn(100))
	}
	fmt.Println(nums)
	fmt.Println(sortArray(nums))
}

func Test_findKthLargest(t *testing.T) {
	nums := []int{2, 1}
	findKthLargest(nums, 1)
}

func Test_reverseKGroup(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h := buildList(nums)
	fmt.Println(convertList(h))
	fmt.Println(convertList(reverseKGroup(h, 3)))
}

func Test_slice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums[1:])
	fmt.Println(nums[2:])
	fmt.Println(nums[3:])
}
