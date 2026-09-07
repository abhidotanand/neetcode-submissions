func wordPattern(pattern string, s string) bool {
	var n int = len(pattern)
    var hm map[byte]string = make(map[byte]string)
	var splitted []string = strings.Split(s, " ")

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