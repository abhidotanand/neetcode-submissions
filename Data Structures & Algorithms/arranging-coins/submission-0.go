func arrangeCoins(n int) int {
	var l, r int = 1, n
	var mid int

	for l <= r {
		mid = (l + r)/2
		summ := mid*(mid+1)/2

		if summ == n {
			return mid
		} else if summ > n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return mid - 1
}
