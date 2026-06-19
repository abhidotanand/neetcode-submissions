func evalRPN(tokens []string) int {
	var stacks []int = make([]int, 0)

	for _, item := range tokens {
		if num, err := strconv.Atoi(item); err == nil {
			stacks = append(stacks, num)
		} else {
			tmp := 0
			n := len(stacks)
			if item == "+" {
				tmp = stacks[n - 2] + stacks[n - 1]
			} else if item == "-" {
				tmp = stacks[n - 2] - stacks[n - 1]
			} else if item == "*" {
				tmp = stacks[n - 2] * stacks[n - 1]
			} else if item == "/" {
				tmp = stacks[n - 2] / stacks[n - 1]
			}
			stacks = stacks[:n - 2]
			stacks = append(stacks, tmp)
		}
	}
	return stacks[0]
}
