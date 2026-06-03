func isAnagram(s string, t string) bool {
	var n int = len(s)
	var m [26]int

	for i := 0; i < n; i++ {
		m[rune(s[i]) - 97]++
	}

	for i := 0; i < n; i++ {
		m[rune(t[i]) - 97]--
	}

	fmt.Println(m)

	for i := 0; i < 26; i++ {
		if m[i] != 0 {
			return false
		}
	}

	return true
}
