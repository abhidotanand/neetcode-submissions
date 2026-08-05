func pivotIndex(nums []int) int {
	var n int = len(nums)
	var left_sum []int = make([]int, n)
	var right_sum []int = make([]int, n)

	for i := 1; i < n; i++ {
		left_sum[i] = left_sum[i-1] + nums[i-1]
	}

	for i := n-2; i >= 0; i-- {
		right_sum[i] = right_sum[i+1] + nums[i+1]
	}

	for i := 0; i < n; i++ {
		if left_sum[i] == right_sum[i] {
			return i
		}
	}
	return -1
}
