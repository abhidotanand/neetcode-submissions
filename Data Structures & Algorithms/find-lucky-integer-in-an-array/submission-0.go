func findLucky(arr []int) int {
	var n int = len(arr)
	var ans int = -1
	var m map[int]int = make(map[int]int)

	for i := 0; i < n; i++ {
		if _, ok := m[arr[i]]; ok {
			m[arr[i]]++
		} else {
			m[arr[i]] = 1
		}
	}

	for i := 0; i < n; i++ {
		val, _ := m[arr[i]]
		if val == arr[i] {
			ans = max(ans, arr[i])
		}
	}
	return ans
}
