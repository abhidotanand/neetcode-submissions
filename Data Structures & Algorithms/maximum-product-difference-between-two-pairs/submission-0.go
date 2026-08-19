import "slices"
func maxProductDifference(nums []int) int {
	var n int = len(nums)
    var mini1 int = slices.Min(nums)
	var maxi1 int = slices.Max(nums)
	var mini2, maxi2 int = 10001, 0

	for i := 0; i < n; i++ {
		if nums[i] > mini1 && nums[i] < mini2 {
			mini2 = nums[i]
		}

		if nums[i] < maxi1 && nums[i] > maxi2 {
			maxi2 = nums[i]
		}
	}

	return maxi2*maxi1 - mini2*mini1
}