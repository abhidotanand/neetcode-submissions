func appendCharacters(s string, t string) int {
    var n, m int = len(s), len(t)
	var i, j int

	for ;i < n;i++ {
		if t[j] == s[i] {
			j++
		}
		if j == m {
			return 0
		}
	}
	return m - j
}