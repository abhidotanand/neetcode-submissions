func countCharacters(words []string, chars string) int {
    var n int = len(chars)
	var m int = len(words)
	var all_chars []byte = make([]byte, 26)
	var tmp []byte = make([]byte, 26)
	var ans int
	var flag bool

	for i := 0; i < n; i++ {
		all_chars[chars[i] - 'a']++
	}

	for i := 0; i < m; i++ {
		copy(tmp, all_chars)
		flag = false
		for j := 0; j < len(words[i]); j++ {
			if tmp[words[i][j] - 'a'] == 0 {
				flag = true
				break
			} else {
				tmp[words[i][j] - 'a']--
			}
		}
		if !flag{
			ans += len(words[i])
		}
	}

	return ans
}