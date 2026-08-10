func maxNumberOfBalloons(text string) int {
	var n int = len(text)
    var freq_arr []int = make([]int, 26)

	for i := 0; i < n; i++ {
		freq_arr[text[i] - 97]++
	}

	return min(freq_arr[0], freq_arr[1], freq_arr[11]/2, freq_arr[14]/2, freq_arr[13])
}