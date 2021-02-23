package daily

func maxSatisfied(customers []int, grumpy []int, X int) int {
	base := 0
	for i := range customers {
		if grumpy[i] == 0 {
			base = base + customers[i]
		}
	}
	add := 0
	for i := 0; i < X && i < len(customers); i++ {
		if grumpy[i] == 1 {
			add = add + customers[i]
		}
	}
	maxAdd := add
	for i := X; i < len(customers); i++ {
		add = add + customers[i]*grumpy[i] - customers[i-X]*grumpy[i-X]
		if add > maxAdd {
			maxAdd = add
		}
	}
	return base + maxAdd
}
