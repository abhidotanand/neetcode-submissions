func sortColors(nums []int) {
    var n int = len(nums)
	var front, rear int = 0, n-1

	for i := 0; i <= rear;{
		if nums[i] == 0 {
			nums[i], nums[front] = nums[front], nums[i]
			front++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[rear] = nums[rear], nums[i]
			rear--
		} else {
			i++
		}
	}
}
