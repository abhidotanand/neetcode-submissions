func divideArray(nums []int) bool {
    var n int = len(nums)
	var m map[int]int = make(map[int]int)

	for i := 0; i < n; i++ {
		m[nums[i]]++
	}

	for i := 0; i < n; i++ {
		val := m[nums[i]]
		if val % 2 == 1 {
			return false
		}
	}

	return true
}