func groupAnagrams(strs []string) [][]string {
	var m map[[26]byte][]string = make(map[[26]byte][]string)
	var buff [][26]byte
	var tmp [26]byte = [26]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	var n int = len(strs)
	var ans [][]string

	for i := 0; i < n; i++ {
		for j := 0; j < len(strs[i]); j++ {
			tmp[strs[i][j] - 'a']++
		}
		if _, ok := m[tmp]; ok {
			m[tmp] = append(m[tmp], strs[i])
		} else {
			m[tmp] = []string{strs[i]}
			buff = append(buff, tmp)
		}
		tmp = [26]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	}

	for i := 0; i < len(buff); i++ {
		ans = append(ans, m[buff[i]])
	}

	return ans
}
