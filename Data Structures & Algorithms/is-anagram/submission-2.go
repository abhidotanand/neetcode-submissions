func isAnagram(s string, t string) bool {
	var sn int = len(s)
	var tn int = len(t)
	var m [26]int

	if sn != tn {
		return false
	}
	
	for i := 0; i < sn; i++ {
		m[rune(s[i]) - 97]++
	}

	for i := 0; i < sn; i++ {
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
