package basic_calculator_ii

func calculate(s string) int {
	n, _ := cal(s, 0)
	return n
}

func cal(s string, start int) (value int, end int) {
	var res, n int
	var lastCal uint8 = '+'
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '+', '-':
			res = sampleCal(res, n, lastCal)
			n = 0
			lastCal = s[i]
		case '*', '/':
			m, skip := readNextNum(s, i+1)
			n = sampleCal(n, m, s[i])
			i = skip - 1
		case ' ':
			continue
		default:
			n = n*10 + int(s[i]-'0')
		}
	}
	return sampleCal(res, n, lastCal), len(s)
}

func readNextNum(s string, start int) (value, end int) {
	for i := start; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			value = value*10 + int(s[i]-'0')
		} else {
			return value, i
		}
	}
	return value, len(s)
}

func sampleCal(a, b int, cal byte) int {
	switch cal {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}
