func removeDuplicates(s string, k int) string {
	var n int = len(s)
	var m map[byte]int = make(map[byte]int)
	var stack []byte = make([]byte, 0)
	var ans string

	for i := 0; i < n; i++ {
		stack = append(stack, s[i])
		m[s[i]]++

		if m[s[i]] >= k {
			for len(stack) > 0 && stack[len(stack)-1] == s[i] {
				stack = stack[:len(stack)-1]
				m[s[i]]--
			}
		}
	}

	for len(stack) > 0 {
		ans = string(stack[len(stack)-1]) + ans
		stack = stack[:len(stack)-1]
	}
	return ans
}
