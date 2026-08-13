func removeDuplicates(s string, k int) string {
	type Pair struct{
		ch byte
		cnt int
	}

	var n int = len(s)
	var stack []Pair = make([]Pair, 0)
	var ans string

	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, Pair{ch: s[i], cnt: 1})
		} else {
			if s[i] == stack[len(stack) - 1].ch {
				stack = append(stack, Pair{ch: s[i], cnt: stack[len(stack) - 1].cnt + 1})
				if stack[len(stack) - 1].cnt >= k {
					for len(stack) > 0 && stack[len(stack) - 1].ch == s[i] {
						stack = stack[:len(stack)-1]
					}
				}

			} else {
				stack = append(stack, Pair{ch: s[i], cnt: 1})
			}
		}
	}
	
	for len(stack) > 0 {
		ans = string(stack[len(stack) - 1].ch) + ans
		stack = stack[:len(stack) - 1]
	}

	return ans
}
