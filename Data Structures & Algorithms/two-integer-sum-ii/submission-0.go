func twoSum(numbers []int, target int) []int {
	var n int = len(numbers)
	var front, rear int = 0, n-1
	var summ int

	for front < rear {
		summ = numbers[front] + numbers[rear]
		if summ > target {
			rear--
		} else if summ < target {
			front++
		} else {
			return []int{front+1, rear+1}
		}
	}
	return []int{}
}
