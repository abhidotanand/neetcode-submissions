func mergeAlternately(word1 string, word2 string) string {
	var n1 int = len(word1)
	var n2 int = len(word2)
	var tmp1, tmp2 int
	var ans string

	for tmp1 < n1 && tmp2 < n2 {
		ans += string(word1[tmp1]) + string(word2[tmp2])
		tmp1++
		tmp2++
	}

	if tmp1 == n1 {
		for i := tmp2; i < n2; i++ {
			ans += string(word2[i])
		}
	} else if tmp2 == n2 {
		for i := tmp1; i < n1; i++ {
			ans += string(word2[i])
		}
	}
	return ans
}
