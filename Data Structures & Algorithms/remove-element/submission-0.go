func removeElement(nums []int, val int) int {
	var n, cnt int = len(nums), 0


    for i := 0; i < n; i++ {
		if nums[i] == val {
			nums[i] = -1
			cnt++
		}
	}

	var front, rear int = 0, n - 1

	for front < rear {
		for rear >= 0 && nums[rear] == -1 {
			rear--
		}

		if rear <= front{
			break
		}

		if nums[front] == -1 {
			nums[front], nums[rear] = nums[rear], nums[front]
			rear--
		}
		front++
	}
	return n - cnt
}
