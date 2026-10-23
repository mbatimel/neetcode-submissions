func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make(map[byte]int, len(s))
	lastVAl := s[0]
	stack[lastVAl] += 1
	for i := 1; i < len(s); i++ {
		c := s[i]
		switch {
		case c == 40:
			if stack[c] < 0 {
				return false
			}
			lastVAl = c
			stack[c] += 1
		case c == 41:
			if lastVAl == 123 || lastVAl == 91 {
				return false
			}
			lastVAl = c
			stack[40] -= 1
		case c == 91:
			if stack[c] < 0 {
				return false
			}
			lastVAl = c
			stack[c] += 1
		case c == 93:
			if lastVAl == 123 || lastVAl == 40 {
				return false
			}
			lastVAl = c
			stack[91] -= 1
		case c == 123:
			if stack[c] < 0 {
				return false
			}
			lastVAl = c
			stack[c] += 1
		case c == 125:
			if lastVAl == 91 || lastVAl == 40 {
				return false
			}
			lastVAl = c
			stack[123] -= 1

		}

	}
	for _, i := range stack {
		if i != 0 {
			return false
		}
	}

	return true

}