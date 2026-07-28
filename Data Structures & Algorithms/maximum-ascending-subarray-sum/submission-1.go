func maxAscendingSum(nums []int) int {
    var ans int = nums[0]
	var n int = len(nums)
	var tmp int = nums[0]

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			tmp += nums[i]
		} else {
			tmp = nums[i]
		}
		ans = max(ans, tmp)
	}
	return ans
}