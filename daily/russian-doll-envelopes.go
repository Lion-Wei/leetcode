package daily

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})

	f := []int{}
	for i := range envelopes {
		if index := sort.SearchInts(f, envelopes[i][1]); index < len(f) {
			f[index] = envelopes[i][1]
		} else {
			f = append(f, envelopes[i][1])
		}
	}
	return len(f)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
