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
	tmp := left

	left = 0
	right = tmp - 1
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

	left = tmp
	right = n-1
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
