func isPalindrome(s string) bool {
	var n int = len(s)
	var front, rear int = 0, n-1

	for front < rear {
		for !isAlphabet(byte(s[front])) {
			front++
			if front >= n{
				break
			}
		}

		for !isAlphabet(byte(s[rear])) {
			rear--
			if rear < 0 {
				break
			}
		}

		fmt.Println(front, rear)

		if front >= rear {
			break
		}
		if strings.ToLower((string(s[front]))) != strings.ToLower((string(s[rear]))) {
			return false
		}
		front++
		rear--
	}
	return true
}

func isAlphabet(i byte) bool {
	return (i >= 65 && i <= 90) || (i >= 97 && i <= 122)
}
