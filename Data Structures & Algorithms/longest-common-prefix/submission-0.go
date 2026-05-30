func longestCommonPrefix(strs []string) string {
	var n = len(strs)
    var ans string
	var m, i, j int = 201, 0, 0

	for i := range strs{
		if len(strs[i]) < m {
			m = len(strs[i])
		}
	}

	for ; j < m; j++{
		for ; i < n-1; i++{
			if strs[i][j] != strs[i+1][j]{
				return ans
			}
		}
		ans += string(strs[i][j])
		i = 0
	}
	return ans
}
