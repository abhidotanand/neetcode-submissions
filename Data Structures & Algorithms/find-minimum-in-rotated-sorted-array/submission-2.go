func findMin(nums []int) int {
	var n int = len(nums)
	var left, right int = 0, n-1
	var ans int = nums[0]

	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return min(nums[0], nums[1])
	}

	for left < right {
		mid := (left+right)/2

		if mid == 0 || mid == n-1 {
			ans = min(ans, nums[mid])
			break
		}
		
		ans = min(ans, nums[mid], nums[mid-1], nums[mid+1])

		if min(nums[mid-1], nums[left]) < min(nums[mid+1], nums[right]) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
