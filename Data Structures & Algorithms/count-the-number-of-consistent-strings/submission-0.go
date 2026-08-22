func countConsistentStrings(allowed string, words []string) int {
    var n, m int = len(allowed), len(words)
	var mp map[byte]struct{} = make(map[byte]struct{})
	var ans int
	var flag bool

	for i := 0; i < n; i++ {
		mp[allowed[i]] = struct{}{}
	}

	for i := 0; i < m; i++ {
		flag =  false
		for j := 0; j < len(words[i]); j++ {
			if _, ok := mp[words[i][j]]; !ok {
				flag = true
				break
			}
		}
		if !flag {
			ans++
		}
	}
	return ans
}