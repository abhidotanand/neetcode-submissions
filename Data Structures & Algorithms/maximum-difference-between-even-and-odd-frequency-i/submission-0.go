func maxDifference(s string) int {
	var n int = len(s)
	var m map[byte]int = make(map[byte]int)
	var odd, even int = -1, 1e9

	for i := range s {
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}

	for i := 0; i < n; i++ {
		if val, _ := m[s[i]]; val%2 == 1 && val > odd {
			odd = val
		}
	}

	for i := 0; i < n; i++ {
		if val, _ := m[s[i]]; val%2 == 0 && val < even {
			even = val
		}
	}

	return odd - even
}
