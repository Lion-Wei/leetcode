package struct_algorithm

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_fastSort(t *testing.T) {
	nums := []int{}
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100))
	}
	fmt.Println(nums)
	fmt.Println(fastSort(nums))
}
