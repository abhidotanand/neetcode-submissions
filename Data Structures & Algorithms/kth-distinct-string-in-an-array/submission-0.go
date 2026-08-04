func kthDistinct(arr []string, k int) string {
	var n int = len(arr)
	var m map[string]int = make(map[string]int)
	var cnt int

	for i := 0; i < n; i++ {
		if _, ok := m[arr[i]]; ok {
			m[arr[i]]++
		} else {
			m[arr[i]] = 1
		}
	}

	for i := 0; i < n; i++ {
		if m[arr[i]] == 1 {
			cnt++
			if cnt == k {
				return arr[i]
			}
		}
	}
	return ""
}