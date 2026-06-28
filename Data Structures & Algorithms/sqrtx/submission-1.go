func mySqrt(x int) int {
	if x == 1 {
		return 1
	} else if x == 0 {
		return 0
	}
	var n int = 2

	for n*n <= x {
		n += 1
	}
	return n-1
}
