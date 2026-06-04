func longestCommonPrefix(strs []string) string {
    var n int = len(strs)
	var m int = 201
	var ans string = ""

	for i := 0; i < n; i++ {
		if len(strs[i]) < m {
			m = len(strs[i])
		}
	}

	var i int
	for j := 0; j < m; j++ {
		for ; i < n - 1; i++ {
			if strs[i][j] != strs[i+1][j] {
				return ans
			}
		}
		i = 0
		ans += string(strs[i][j])
	}
	return ans
}
