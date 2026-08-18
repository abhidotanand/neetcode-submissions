func numIdenticalPairs(nums []int) int {
    var ans int
	var n int = len(nums)
	var m map[int]int = make(map[int]int)

	for i := 0; i < n; i++ {
		m[nums[i]]++
	}

	for i := 0; i < n; i++ {
		if cnt, ok := m[nums[i]]; ok {
			ans += cnt*(cnt-1)/2
			delete(m, nums[i])
		}
	}
	return ans
}