func lengthOfLongestSubstring(s string) int {
	var n int = len(s)
	var front, rear int
	var ans int
	var m map[byte]int = make(map[byte]int)

	for rear < n {
		if _, ok := m[s[rear]]; !ok {
			m[s[rear]] = rear
			rear++
		} else {
			front = m[s[rear]] + 1
			rear = front
			clear(m)
		}
		ans = max(ans, rear - front)
	}
	return ans
}
