func longestConsecutive(nums []int) int {
	var n int = len(nums)
	var m map[int]struct{} = make(map[int]struct{})
	var tmp, curr, ans int

	for i := 0; i < n; i++ {
		m[nums[i]] = struct{}{}
	}

	for i := 0; i < n; i++ {
		tmp = 0
		
		curr = nums[i]
		for _, ok := m[curr]; ok; {
			tmp += 1
			delete(m, curr)
			curr--
			_, ok = m[curr]
		}

		curr = nums[i]+1
		for _, ok := m[curr]; ok;  {
			tmp += 1
			delete(m, curr)
			curr++
			_, ok = m[curr]
		}

		ans = max(ans, tmp)
	}
	return ans
}
