func sortArrayByParity(nums []int) []int {
	var n int = len(nums)
	var front, rear int = 0, n-1

	for front < rear {
		for front < rear && nums[front] % 2 == 0 {
			front++
		}
		for front < rear && nums[rear] % 2 == 1 {
			rear--
		}

		if front < rear {
			nums[front], nums[rear] = nums[rear], nums[front]
		}
		front++
		rear--
	}

	return nums
}
