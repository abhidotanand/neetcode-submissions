func isAnagram(s string, t string) bool {
	var m int = len(s)
	var n int = len(t)

	if m != n {
		return false
	}

	var sa [26]int
	var ta [26]int

	for i := 0; i < m; i++ {
		sa[s[i] - 'a']++
	}

	for i := 0; i < n; i++ {
		ta[t[i] - 'a']++
	}

	if sa == ta {
		return true
	}
	return false
}
