func validPalindrome(s string) bool {
	var n int = len(s)
	var front, rear int = 0, n-1
	var flag1, flag2 bool = true, true
	var tmp1_front, tmp1_rear, tmp2_front, tmp2_rear int

	for front < rear {
		if s[front] != s[rear] {
			tmp1_front = front+1
			tmp1_rear = rear
			for tmp1_front < tmp1_rear {
				if s[tmp1_front] != s[tmp1_rear] {
					flag1 = false
					break
				}
				tmp1_rear--
				tmp1_front++
			}

			tmp2_front = front
			tmp2_rear = rear-1
			for tmp2_front < tmp2_rear {
				if s[tmp2_front] != s[tmp2_rear] {
					flag2 = false
					break
				}
				tmp2_front++
				tmp2_rear--
			}
			return flag1 || flag2
		}
		front++
		rear--
	}
	return true
}
