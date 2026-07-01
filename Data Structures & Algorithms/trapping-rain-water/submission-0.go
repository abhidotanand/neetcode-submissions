func trap(height []int) int {
	var n int = len(height)
	var left, right = 0, n-1
	var left_max, right_max int = height[left], height[right]
	var res int

	for left < right {
		if left_max < right_max {
			left++
			res += max(0, left_max - height[left])
			left_max = max(left_max, height[left])
		} else {
			right--
			res += max(0, right_max - height[right])
			right_max = max(right_max, height[right])
		}
	}
	return res
}