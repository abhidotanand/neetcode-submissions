func firstMissingPositive(nums []int) int {
	var n int = len(nums)

	for i := 0; i < n; i++ {
		if nums[i] > n || nums[i] < 0 {
			nums[i] = 0
		}
	}

	for i := 0; i < n; i++ {
		for (nums[i] != i+1) && (nums[i] != 0) && (nums[i] != nums[nums[i]-1]) {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i+1
		}
	}
	return n+1
}
