func lengthOfLastWord(s string) int {
	var n int = len(s)
	var start bool
	var cnt int

	for i := n-1; i >= 0; i-- {
		if string(s[i]) != " " {
			start = true
			cnt++
		} else if string(s[i]) == " " && start{
			return cnt
		}
	}
	return cnt
}
