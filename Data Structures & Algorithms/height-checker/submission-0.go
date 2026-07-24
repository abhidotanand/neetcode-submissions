import "slices"
func heightChecker(heights []int) int {
    var n int = len(heights)
	var cnt int
	var tmp []int = make([]int, n)

	_ = copy(tmp, heights)
	slices.Sort(tmp)

	for i := 0; i < n; i++ {
		if heights[i] != tmp[i] {
			cnt++
		}
	}
	return cnt
}