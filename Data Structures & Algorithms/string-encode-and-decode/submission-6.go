
type Solution struct {
	code []string
}

func (s *Solution) Encode(strs []string) string {
	s.code = strs
	var res string
	for i := 0; i < len(strs); i++ {
		res += strs[i]
		if i != len(strs)-1 {
			res += ""
		}

	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	for i := range s.code {
		res = append(res, s.code[i])
	}
	return res
}