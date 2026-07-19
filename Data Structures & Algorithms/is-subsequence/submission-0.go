func isSubsequence(s string, t string) bool {
	var n, m int = len(s), len(t)
	var i, j int

	for j = 0; j < m; j++ {
		if t[j] == s[i] {
			i++
		}
		if i == n {
			return true
		}
	}
	return false
}