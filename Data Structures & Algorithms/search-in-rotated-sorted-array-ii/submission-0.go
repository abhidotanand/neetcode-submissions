func search(nums []int, target int) bool {
	var n int = len(nums)
	var left, right int = 0, n-1

	for left < right {
		mid := (left+right)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if target < nums[left] {
		right = left - 1
		left = 0
	}

	for left <= right {
		mid := (left+right)/2

		if nums[mid] == target {
			return true
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
