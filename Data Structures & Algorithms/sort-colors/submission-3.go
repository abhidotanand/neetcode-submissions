func sortColors(nums []int) {
    var n int = len(nums)
	var i, front, rear int = 0, 0, n-1

	for i <= rear {
		if nums[i] == 0 {
			nums[front], nums[i] =  nums[i], nums[front]
			front++
			if nums[i] == 1 {
				i++
			}
		} else if nums[i] == 2 {
			nums[rear], nums[i] = nums[i], nums[rear]
			rear--
			if nums[i] == 1 {
				i++
			}
		} else {
			i++
		}
	}
}