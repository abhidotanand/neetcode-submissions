func asteroidCollision(asteroids []int) []int {
	var stack []int = make([]int, 0)

	for _, ast := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, ast)
		} else {
			if ast < 0 {
				if stack[len(stack)-1] > 0{
					if -ast > stack[len(stack)-1] {
						stack = stack[:len(stack)-1]
						for len(stack) > 0 {
							if -ast > stack[len(stack)-1] {
								stack = stack[:len(stack)-1]
							} else if -ast == stack[len(stack)-1]{
								stack = stack[:len(stack)-1]
							} else {
								break
							}
						}
						if len(stack) == 0 {
							stack = append(stack, ast)
						}
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
