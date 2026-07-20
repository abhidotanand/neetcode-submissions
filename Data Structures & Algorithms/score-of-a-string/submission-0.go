func scoreOfString(s string) int {
	var n int = len(s)
	var ans int

	for i := 1; i < n; i++ {
		ans += int(math.Abs(float64(s[i])-float64(s[i-1])))
	}

	return ans
}
