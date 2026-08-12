func findDisappearedNumbers(nums []int) []int {
	var n int = len(nums)
	var ans []int

	for i := 0; i < n; i++ {
		if nums[int(math.Abs(float64(nums[i])))-1] > 0 {
			nums[int(math.Abs(float64(nums[i])))-1] = -nums[int(math.Abs(float64(nums[i])))-1]
		}
	}
	
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}