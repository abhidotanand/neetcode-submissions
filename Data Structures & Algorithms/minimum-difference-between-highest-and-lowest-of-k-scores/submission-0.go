import "slices"
func minimumDifference(nums []int, k int) int {
	slices.Sort(nums)
	var front, rear int = 0, k-1
	var n int = len(nums)
	var ans int = 100001

	for rear < n {
		ans = min(ans, nums[rear] - nums[front])
		front++
		rear++
	}

	return ans
}
