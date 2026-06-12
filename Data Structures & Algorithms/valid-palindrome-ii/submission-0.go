func validPalindrome(s string) bool {
	var n int = len(s)
	var front, rear int = 0, n-1
	var flag bool

	for front < rear {
		if s[front] != s[rear] {
			if flag {
				return false
			} else {
				flag = true
				if s[front] == s[rear-1]{
					rear--
				} else if s[front+1] == s[rear]{
					front++
				} else {
					return false
				}
				continue
			}
		}
		front++
		rear--
	}
	return true
}
