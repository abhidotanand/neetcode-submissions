func rotate(nums []int, k int) {
	reverser(nums[:k])
	reverser(nums[k:])
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
