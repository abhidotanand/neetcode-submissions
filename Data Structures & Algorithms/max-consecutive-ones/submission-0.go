func findMaxConsecutiveOnes(nums []int) int {
	var n int = len(nums)
	var ans, tmp int

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			tmp = 0
		} else {
			tmp++
		}
		ans = max(tmp, ans)
	}
	return ans
}
