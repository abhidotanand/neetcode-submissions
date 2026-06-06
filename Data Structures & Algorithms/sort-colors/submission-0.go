func sortColors(nums []int) {
	var n int = len(nums)
    var front, rear int = 0, n - 1

	for front < n {
		if nums[front] == 0 {
			front++
		} else {
			break
		}
	}

	for rear > 0 {
		if nums[rear] == 2 {
			rear--
		} else {
			break
		}
	}

	if front >= rear{
		return
	}

	for i := front; i <= rear; i++ {
		if nums[i] == 0 {
			nums[i], nums[front] = nums[front], nums[i]
			front++
		} else if nums[i] == 2 {
			nums[i], nums[rear] = nums[rear], nums[i]
			rear--
		}
	}
}
