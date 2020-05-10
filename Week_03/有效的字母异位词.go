func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_b := []byte(s)
	s_t := []byte(t)
	sort.Slice(s_b, func(i, j int) bool {
		return s_b[i] > s_b[j]
	})
	sort.Slice(s_t, func(i, j int) bool {
		return s_t[i] > s_t[j]
	})
	for i := 0; i < len(s_b); i++ {
		if s_b[i] != s_t[i] {
			return false
		}
	}
	return true
}