func longestMonotonicSubarray(nums []int) int {
    var ans, tmp int = 1, 1
	var n int = len(nums)

	for i := 0; i < n-1; i++ {
		if nums[i+1] > nums[i] {
			tmp++
		} else {
			tmp = 1
		}
		ans = max(ans, tmp)
	}

	tmp = 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] < nums[i] {
			tmp++
		} else {
			tmp = 1
		}
		ans = max(ans, tmp)
	}
	return ans
}