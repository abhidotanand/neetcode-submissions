func wordPattern(pattern string, s string) bool {
    var hm map[byte]string = make(map[byte]string)
	var splitted []string = strings.Split(s, " ")
	var n int = len(pattern)
	var m int = len(splitted)

	if n != m {
		return false
	}

	for i := 0; i < n; i++ {
		if val, ok := hm[pattern[i]]; ok {
			if val != splitted[i] {
				return false
			}
		} else {
			hm[pattern[i]] = splitted[i]
		}
	}

	return true
}