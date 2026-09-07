func wordPattern(pattern string, s string) bool {
    var hm1 map[byte]string = make(map[byte]string)
	var hm2 map[string]byte = make(map[string]byte)
	var splitted []string = strings.Split(s, " ")
	var n int = len(pattern)
	var m int = len(splitted)

	if n != m {
		return false
	}

	for i := 0; i < n; i++ {
		if val, ok := hm1[pattern[i]]; ok {
			if val != splitted[i] {
				return false
			}
		} else {
			hm1[pattern[i]] = splitted[i]
		}
	}

	for i := 0; i < n; i++ {
		if val, ok := hm2[splitted[i]]; ok {
			if val != pattern[i] {
				return false
			}
		} else {
			hm2[splitted[i]] = pattern[i]
		}
	}

	return true
}