import "slices"
func numRescueBoats(people []int, limit int) int {
	var n int = len(people)
	var front, rear = 0, n-1
	var ans int

	slices.Sort(people)

	for front <= rear {
		if people[front] + people[rear] > limit {
			rear--
		} else if people[front] + people[rear] < limit {
			front++
		} else {
			front++
			rear--
		}
		ans++
	}
	return ans
}
