func search(nums []int, target int) int {
	var n int = len(nums)
	var left, right int = 0, n-1

	for left <= right {
		mid := (left+right)/2

		if nums[mid] == target {
			return mid
		}

		if mid == 0 || mid == n-1 {
			return -1
		}

		if nums[left] <= nums[mid-1] && nums[mid+1] >= nums[right] {
			if target <= nums[mid-1] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] >= nums[mid-1] && nums[mid+1] <= nums[right] {
			if target >= nums[mid+1] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[left] && target <= nums[mid-1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
