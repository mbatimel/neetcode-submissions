func isPalindrome(s string) bool {
	s1Lower := strings.ToLower(s)
	buf := bytes.NewBuffer(make([]byte, 0, len(s)))

	for i := 0; i < len(s1Lower); i++ {

		b := s1Lower[i]
		if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
			buf.WriteByte(b)
		}
	}
	s = buf.String()
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true

}