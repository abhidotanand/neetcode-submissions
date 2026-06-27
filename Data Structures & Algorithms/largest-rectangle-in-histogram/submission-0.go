func largestRectangleArea(heights []int) int {
	var n int = len(heights)
	var stack []int = make([]int, 0)
	var lmost []int = make([]int, n)
	var rmost []int = make([]int, n)
	var ans int

	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			lmost[i] = -1
		} else {
			if heights[i] > heights[stack[len(stack)-1]] {
				lmost[i] = stack[len(stack)-1]
				stack = append(stack, i)
			} else {
				for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
					stack = stack[:len(stack)-1]
				}
				if len(stack) > 0 {
					lmost[i] = stack[len(stack)-1]
				} else {
					stack = append(stack, i)
					lmost[i] = -1
				}
			}
		}
	}
	
	stack = stack[:0]
	for i := n-1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, i)
			rmost[i] = n
		} else {
			if heights[i] > heights[stack[len(stack)-1]] {
				rmost[i] = stack[len(stack)-1]
				stack = append(stack, i)
			} else {
				for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
					stack = stack[:len(stack)-1]
				}
				if len(stack) > 0 {
					rmost[i] = stack[len(stack)-1]
				} else {
					stack = append(stack, i)
					rmost[i] = n
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(rmost[i]-i-1) + heights[i]*(i-lmost[i]-1) +  heights[i])
	}

	return ans
}
