func twoSum(nums []int, target int) []int {
	var n int = len(nums)
    var m map[int]int = make(map[int]int)

	for i := 0; i < n; i++ {
		m[nums[i]] = i
	}

	for i := 0; i < n; i++ {
		if _, ok := m[target - nums[i]]; ok && i != m[target - nums[i]] {
			return []int{i, m[target - nums[i]]}
		}
	}
	return []int{-1, -1}
}
