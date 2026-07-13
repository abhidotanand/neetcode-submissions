import "slices"
func checkInclusion(s1 string, s2 string) bool {
	var n1 int = len(s1)
	var n2 int = len(s2)
	var freq_arr []int = make([]int, 26)
	var window []int = make([]int, 26)

	for i := range s1 {
		freq_arr[s1[i] - 97]++
	}

	for i := 0; i < n1; i++ {
		window[s2[i] - 97]++
	}

	if slices.Equal(freq_arr, window) {
		return true
	}

	front, rear := 0, n1
	for ;rear < n2; front, rear = front+1, rear+1 {
		window[s2[rear] - 97]++
		window[s2[front] - 97]--
		if slices.Equal(freq_arr, window) {
			return true
		}
	}
	return false
}
