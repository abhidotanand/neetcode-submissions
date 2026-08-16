func isMonotonic(nums []int) bool {
    var n int = len(nums)
	var dec bool

	if n == 1 {
		return true
	}

	if nums[0] > nums[n-1] {
		dec = true
	}

	if dec {
		for i := 0; i < n-1; i++ {
			if nums[i+1] > nums[i] {
				return false
			}
		}
	} else {
		for i := 0; i < n-1; i++ {
			if nums[i+1] < nums[i] {
				return false
			}
		}
	}

	return true
}