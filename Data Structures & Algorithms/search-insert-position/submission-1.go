func searchInsert(nums []int, target int) int {
	var n int = len(nums)
	var left, right int = 0, n-1
	var mid int

	for left <= right {
		mid = (left + right)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if mid == 0 {
		return mid
	}
	return mid+1
}
