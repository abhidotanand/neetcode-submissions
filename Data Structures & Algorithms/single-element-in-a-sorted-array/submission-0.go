func singleNonDuplicate(nums []int) int {
	var n int = len(nums)
	var left, right int = 1, n-2

	if nums[0] != nums[1] {
		return nums[0]
	}

	if nums[n-1] != nums[n-2] {
		return nums[n-1]
	}

	for left <= right {
		mid := (left+right)/2

		if mid%2 == 1 {

		}
		if nums[mid] == nums[mid-1] {
			if mid%2 == 1 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] == nums[mid+1] {
			if mid%2 == 1 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			return nums[mid]
		}
	}
	return -1
}

// 3 3 4 5 5 6 6 8 8
// 3 3 4 4 5 5 6 8 8

// 3 3 4 5 5 6 6 7 7 10 10
// 3 3 4 4 5 5 6 6 7 10 10