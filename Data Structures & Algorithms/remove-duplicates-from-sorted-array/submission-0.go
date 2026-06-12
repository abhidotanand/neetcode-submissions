func removeDuplicates(nums []int) int {
	var k int = 1
	var n int = len(nums)
	var front int

	for i := 1; i < n; i++{
		if nums[i] != nums[i-1] {
			k++
		} else {
			nums[i-1] = -101
		}
	}

	for front < n && nums[front] != -101 {
		front++
	}

	for i := front+1; i < n; i++ {
		if nums[i] != -101 {
			nums[front] = nums[i]
			front++
		}
	}
	return k
}
