func isIsomorphic(s string, t string) bool {
	var n int = len(s)
	var m map[byte]byte = make(map[byte]byte)

	for i := 0; i < n; i++ {
		if val, ok := m[s[i]]; !ok {
			m[s[i]] = t[i]
		} else {
			if val != t[i] {
				return false
			}
		}
	}
	return true
}
