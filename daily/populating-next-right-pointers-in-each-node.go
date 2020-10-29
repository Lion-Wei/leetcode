package daily

func connect(root *Node) *Node {
	var (
		nexts = []*Node{}
		step  func(n *Node)
		l     int
	)

	step = func(n *Node) {
		if n == nil {
			return
		}
		if len(nexts) > l {
			n.Next = nexts[l]
			nexts[l] = n
		} else {
			nexts = append(nexts, n)
		}
		l++
		step(n.Right)
		step(n.Left)
		l--
	}
	step(root)
	return root
}
