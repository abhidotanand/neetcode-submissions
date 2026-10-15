func sortColors(nums []int) {
    var n int = len(nums)
	var i, front, rear int = 0, 0, n-1

	for i <= rear {
		if nums[i] == 0 {
			if i >= front {
				nums[front], nums[i] =  nums[i], nums[front]
			} else {
				i++
			}
			front++
		} else if nums[i] == 2 {
			nums[rear], nums[i] = nums[i], nums[rear]
			rear--
		} else {
			i++
		}
	}
}