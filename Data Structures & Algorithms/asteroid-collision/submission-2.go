func asteroidCollision(asteroids []int) []int {
	var stack []int = make([]int, 0)

	for i := 0; i < len(asteroids); i++ {
		ast := asteroids[i]
		if len(stack) == 0 {
			stack = append(stack, ast)
		} else {
			if ast < 0 {
				if stack[len(stack)-1] > 0{
					if -ast > stack[len(stack)-1] {
						stack = stack[:len(stack)-1]
						i--
					} else if -ast == stack[len(stack)-1] {
						stack = stack[:len(stack)-1]
					} else {
						continue
					}
				} else {
					stack = append(stack, ast)
				}
			} else {
				stack = append(stack, ast)
			}
		}
	}
	return stack
}
