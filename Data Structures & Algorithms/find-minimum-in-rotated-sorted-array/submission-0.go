func findMin(nums []int) int {
	var n int = len(nums)
	var left, right int = 0, n-1
	var ans int = 1001

	for left < right {
		mid := (left+right)/2

		ans = min(ans, nums[mid])

		if min(nums[mid-1], nums[left]) < min(nums[mid+1], nums[right]) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
