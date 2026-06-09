func longestConsecutive(nums []int) int {
	var n int = len(nums)
	var m map[int]struct{} = make(map[int]struct{})
	var ans int

	for i := 0; i < n; i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = struct{}{}
		}
	}

	for i := 0; i < n; i++ {
		if _, ok := m[nums[i]]; ok {
			cnt := 0
			tmp := nums[i]
			for _, ok := m[tmp]; ok; {
				cnt++
				delete(m, tmp)
				tmp++
				_, ok = m[tmp]
			}

			tmp = nums[i] - 1

			for _, ok := m[tmp]; ok; {
				cnt++
				delete(m, tmp)
				tmp--
				_, ok = m[tmp]
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}
