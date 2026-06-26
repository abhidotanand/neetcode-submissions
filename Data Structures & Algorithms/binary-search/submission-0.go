func search(nums []int, target int) int {
	var n int = len(nums)
	var front, rear int = 0, n-1

	for front <= rear {
		mid := (front + rear)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			front = mid + 1
		} else {
			rear = mid - 1
		}
	}
	return -1
}
