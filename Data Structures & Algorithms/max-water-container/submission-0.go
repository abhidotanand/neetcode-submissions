func maxArea(heights []int) int {
	var n int = len(heights)
	var front, rear = 0, n-1
	var ans int

	for front < rear {
		ans = max(ans, min(heights[front], heights[rear]) * (rear - front))
		
		if heights[front] > heights[rear] {
			rear--
		} else {
			front++
		}
	}
	return ans
}
