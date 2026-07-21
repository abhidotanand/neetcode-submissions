import "slices"
func isIsomorphic(s string, t string) bool {
	var n int = len(s)
	var s_freq, t_freq []byte = make([]byte, 26), make([]byte, 26)

	for i := 0; i < n; i++ {
		s_freq[s[i] - 97]++
	}
	
	for i := 0; i < n; i++ {
		t_freq[t[i] - 97]++
	}

	slices.Sort(s_freq)
	slices.Sort(t_freq)

	if slices.Equal(s_freq, t_freq) {
		return true
	}
	return false
}
