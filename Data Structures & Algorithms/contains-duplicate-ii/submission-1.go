func containsNearbyDuplicate(nums []int, k int) bool {
	var n int = len(nums)
	var m map[int]struct{} = make(map[int]struct{})
	var front, rear int

	if k > n-1 {
		k = n-1
	}

	for rear <= k {
		if _, ok := m[nums[rear]]; ok {
			return true
		}
		m[nums[rear]] = struct{}{}
		rear++
	}

	for rear < n {
		delete(m, nums[front])
		if _, ok := m[nums[rear]]; ok {
			return true
		}
		m[nums[rear]] = struct{}{}
		front++
		rear++
	}
	return false
}
