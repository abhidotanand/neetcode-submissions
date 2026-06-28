func mySqrt(x int) int {
	var n int = 2

	for n*n <= x {
		n += 1
	}
	return n-1
}
