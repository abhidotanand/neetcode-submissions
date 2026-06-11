func reverseString(s []byte) {
	var n int = len(s)
	var front, rear int = 0, n-1

	for front < rear {
		s[front], s[rear] = s[rear], s[front]
		front++
		rear--
	}
}
