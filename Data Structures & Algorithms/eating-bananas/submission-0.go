import "slices"
func minEatingSpeed(piles []int, h int) int {
	var n int = len(piles)
	var left int = 1
	var right int = slices.Max(piles)
	var ans int

	for left <= right {
		speed := (left + right)/2
		total := 0
		for i := 0; i < n; i++ {
			total += int(math.Ceil(float64(piles[i])/float64(speed)))
		}
		if total <= h {
			ans = speed
			right = speed - 1
		} else {
			left = speed + 1
		}
	}
	return ans
}
