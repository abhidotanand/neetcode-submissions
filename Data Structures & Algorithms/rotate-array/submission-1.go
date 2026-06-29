func rotate(nums []int, k int) {
	var n int = len(nums)
	k = k%n

	reverser(nums[:n-k])
	reverser(nums[n-k:])
	reverser(nums)
}

func reverser(arr []int) {
	var n int = len(arr)
	var front, rear = 0, n-1
	
	for front < rear {
		arr[front], arr[rear] = arr[rear], arr[front]
		front++
		rear--
	}
}
