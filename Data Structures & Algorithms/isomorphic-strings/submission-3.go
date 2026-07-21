import "slices"
func isIsomorphic(s string, t string) bool {
	var n int = len(s)
	var s_freq []byte = make([]byte, 255)

	for i := 0; i < n; i++ {
		if s_freq[s[i]] == 0 {
			if slices.Contains(s_freq, t[i]) && slices.Index(s_freq, t[i]) != i {
				return false
			}
			s_freq[s[i]] = t[i]
		} else {
			if s_freq[s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}
