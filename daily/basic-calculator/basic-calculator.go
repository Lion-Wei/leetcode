package basic_calculator

func calculate(s string) int {
	n, _ := cal("("+s+")", 0)
	return n
}

func cal(s string, start int) (value int, end int) {
	if s[start] != '(' {
		return 0, 0
	}
	var res, n int
	var lastCal uint8 = '+'
	for i := start + 1; i < len(s); i++ {
		switch s[i] {
		case '+', '-':
			res = sampleCal(res, n, lastCal)
			n = 0
			lastCal = s[i]
		case '(':
			n, skip := cal(s, i)
			res = sampleCal(res, n, lastCal)
			n, i = 0, skip
		case ')':
			res = sampleCal(res, n, lastCal)
			return res, i
		case ' ':
			continue
		default:
			n = n*10 + int(s[i]-'0')
		}
	}
	return sampleCal(res, n, lastCal), len(s)
}

func sampleCal(a, b int, cal byte) int {
	switch cal {
	case '+':
		return a + b
	case '-':
		return a - b
	}
	return 0
}
