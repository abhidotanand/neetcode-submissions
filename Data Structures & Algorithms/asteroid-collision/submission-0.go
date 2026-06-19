func asteroidCollision(asteroids []int) []int {
	var stack []int = make([]int, 0)

	for _, ast := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, ast)
		} else {
			if stack[len(stack) - 1] > 0 && ast < 0 {
				if stack[len(stack) - 1] > -ast{
					continue
				} else if stack[len(stack) - 1] < -ast {
					stack = stack[:len(stack) - 1]
					stack = append(stack, ast)
				} else {
					stack = stack[:len(stack) - 1]
				}
			} else {
				stack = append(stack, ast)
			}
		}
	}
	return stack
}
