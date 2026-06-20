func dailyTemperatures(temperatures []int) []int {
	var n int = len(temperatures)
	var result []int = make([]int, n)
	var stack [][]int

	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, []int{temperatures[i], i})
		} else {
			if stack[len(stack)-1][0] >= temperatures[i] {
				stack = append(stack, []int{temperatures[i], i})
			} else {
				for len(stack) > 0 && stack[len(stack)-1][0] < temperatures[i] {
					result[stack[len(stack)-1][1]] = i - stack[len(stack)-1][1]
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, []int{temperatures[i], i})
			}
		}
	}
	return result
}
