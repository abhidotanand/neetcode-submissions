func calPoints(operations []string) int {
	var ans []int = make([]int, 0)
	var sum int

	for _, o := range operations {
		if o == "+" {
			ans = append(ans, ans[len(ans)-1] + ans[len(ans)-2])
		} else if o == "D" {
			ans = append(ans, ans[len(ans)-1] * 2)
		} else if o == "C" {
			ans = ans[:len(ans)-1]
		} else {
			num, _ := strconv.Atoi(o)
			ans = append(ans, num)
		}
	}

	for i := 0; i< len(ans); i++{
		sum += ans[i]
	}

	return sum
}
