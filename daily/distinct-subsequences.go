package daily

func numDistinct(s string, t string) int {
	distincts := make([]int, len(t)+1)
	distincts[0] = 1
	for i := range s {
		for j := len(t) - 1; j >= 0; j-- {
			if t[j] == s[i] {
				distincts[j+1] = distincts[j+1] + distincts[j]
			}
		}
	}
	return distincts[len(t)]
}
