func largestGoodInteger(num string) string {
	var n int = len(num)
	var rear int = 2
	var ans int
	var flag bool

	for rear < n {
		a, _ := strconv.Atoi(string(num[rear-2]))
		b, _ := strconv.Atoi(string(num[rear-1]))
		c, _ := strconv.Atoi(string(num[rear]))

		tmp := a*100 + b*10 + c
		if num[rear] == num[rear-1] && num[rear] == num[rear-2] && tmp >= ans {
			ans = tmp
			flag = true
		}
		rear++
	}

	if flag {
		return strconv.Itoa(ans/100) + strconv.Itoa(ans/100) + strconv.Itoa(ans/100)
	}
	return ""
}
