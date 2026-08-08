func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var n int = len(customers)
	var ans int
	var total int
	var front, rear int
	var tmp int

	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}

	for rear < minutes {
		if grumpy[rear] == 1 {
			tmp += customers[rear]
		}
		rear++
	}

	total += tmp
	ans = total

	for rear < n {
		if grumpy[rear] == 1 {
			total += customers[rear]
		}
		if grumpy[front] == 1 {
			total -= customers[front]
		}
		ans = max(ans, total)
		front++
		rear++
	}
	return ans
}
