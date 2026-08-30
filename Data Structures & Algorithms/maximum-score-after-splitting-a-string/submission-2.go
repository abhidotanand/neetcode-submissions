func maxScore(s string) int {
    var n int = len(s)
	var total_ones int
	var tmp_ones, tmp_zeroes int
	var ans int

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			total_ones++
		}
	}

	for i := 0; i < n-1; i++ {
		if s[i] == '0' {
			tmp_zeroes++
		} else {
			tmp_ones++
		}

		ans = max(ans, tmp_zeroes + total_ones - tmp_ones)
	}

	return ans
}