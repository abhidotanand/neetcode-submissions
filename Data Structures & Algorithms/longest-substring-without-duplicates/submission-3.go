func lengthOfLongestSubstring(s string) int {
	var n int = len(s)
	var front, rear int
	var m map[byte]int = make(map[byte]int)
	var ans int

	for rear < n {
		if ind, ok := m[s[rear]]; !ok || (ind < front) {
			m[s[rear]] = rear
		} else {
			front = m[s[rear]] + 1
			m[s[rear]] = rear
		}
		ans = max(ans, rear - front + 1)
		rear++
	}
	return ans
}
