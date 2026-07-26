func canConstruct(ransomNote string, magazine string) bool {
	var n int = len(magazine)
	var n2 int = len(ransomNote)
	var m map[byte]int = make(map[byte]int)

	for i := 0; i < n; i++ {
		if _, ok := m[magazine[i]]; ok {
			m[magazine[i]]++
		} else {
			m[magazine[i]] = 1
		}
	}

	for i := 0; i < n2; i++ {
		if _, ok := m[ransomNote[i]]; ok {
			m[ransomNote[i]]--

			if m[ransomNote[i]] < 0 {
				return false
			}
		}
	}
	return true
}
