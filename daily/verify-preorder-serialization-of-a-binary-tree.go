package daily

func isValidSerialization(preorder string) bool {
	emptyPos := 1
	for i := 0; i < len(preorder); i++ {
		switch preorder[i] {
		case ',':
			break
		case '#':
			emptyPos--
		default:
			for ; i < len(preorder) && preorder[i] != ','; i++ {
			}
			emptyPos++
		}
		if emptyPos == 0 {
			return false
		}
	}
	return emptyPos == 0
}
