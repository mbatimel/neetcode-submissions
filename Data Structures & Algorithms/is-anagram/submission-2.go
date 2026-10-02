func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	maps := make(map[byte]int, len(s))

	mapt := make(map[byte]int, len(t))
	sLower := strings.ToLower(s)
	tLower := strings.ToLower(t)
	for i := 0; i < len(s); i++ {
		maps[sLower[i]] += 1
		mapt[tLower[i]] += 1
	}
	for n, _ := range maps {
		if _, ok := mapt[n]; !ok {
			return false
		}
		if mapt[n] != maps[n] {
			return false
		}
	}
	return true

}