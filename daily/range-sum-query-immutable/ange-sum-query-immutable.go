package range_sum_query_immutable

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	res := NumArray{
		sums: make([]int, len(nums)+1),
	}
	for i := range nums {
		res.sums[i+1] = nums[i] + res.sums[i]
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sums[j+1] - this.sums[i]
}
