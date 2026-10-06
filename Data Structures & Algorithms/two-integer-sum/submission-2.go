func twoSum(nums []int, target int) []int {
    var m map[int]int = make(map[int]int)
	var n int = len(nums)

	for i := 0; i < n; i++ {
		m[nums[i]] = i
	}

	for i := 0; i < n; i++ {
		if ind, ok := m[target - nums[i]]; ok && ind != i {
			return []int{i, ind}
		}
	}
	return []int{}
}
