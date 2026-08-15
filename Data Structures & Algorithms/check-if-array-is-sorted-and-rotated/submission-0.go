func check(nums []int) bool {
    var n int = len(nums)
	var max_idx int

	for i := 1; i < n; i++ {
		if nums[i] > nums[max_idx] {
			max_idx = i
		}
	}

	for i := 0; i < max_idx; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	for i := max_idx+1; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	if max_idx == n-1 {
		return true
	} else if nums[0] >= nums[n-1] {
		return true
	}
	return false
}