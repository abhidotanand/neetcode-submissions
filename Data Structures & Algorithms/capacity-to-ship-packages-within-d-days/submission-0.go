import "slices"
func shipWithinDays(weights []int, days int) int {
	var n int = len(weights)
	var min_cap int = slices.Max(weights)
	var max_cap int
	var ans int

	for i := range weights {
		max_cap += weights[i]
	}

	for min_cap <= max_cap {
		mid_cap := (min_cap + max_cap)/2
		
		cnt := 0
		i := 0
		summ := 0
		for i < n {
			if summ + weights[i] <= mid_cap {
				summ += weights[i]
				i++
				continue
			}
			cnt++
			summ = 0
		}
		cnt++

		if cnt <= days {
			ans = mid_cap
			max_cap = mid_cap - 1
		} else {
			min_cap = mid_cap + 1
		}
	}
	return ans
}