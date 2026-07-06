func lengthOfLongestSubstring(s string) int {
	var n int = len(s)
	var front, rear int
	var m map[byte]int = make(map[byte]int)
	var ans int

	for rear < n {
		if _,ok := m[s[rear]]; !ok {
			ans++
			m[s[rear]] = rear
		} else {
			break
		}
		rear++
	}

	tmp := 0
	for rear < n {
		if _, ok := m[s[rear]]; !ok {
			tmp = rear - front + 1
		} else {
			front = m[s[rear]] + 1
			m[s[rear]] = rear
		}
		ans = max(ans, tmp)
		rear++
	}
	return ans
}
