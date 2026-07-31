func isPerfectSquare(num int) bool {
	var l, r = 1, num

	for l <= r {
		mid := (l + r)/2
		prod := mid * mid

		if prod == num {
			return true
		} else if prod > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
