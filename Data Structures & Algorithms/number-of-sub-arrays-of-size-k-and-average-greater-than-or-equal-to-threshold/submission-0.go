func numOfSubarrays(arr []int, k int, threshold int) int {
	var n int = len(arr)
	var front, rear int
	var sum int
	var ans int

	for rear < k {
		sum += arr[rear]
		rear++
	}

	if sum/k >= threshold {
		ans++
	}

	for rear < n {
		sum += arr[rear]
		sum -= arr[front]
		if sum/3 >= threshold {
			ans++
		}
		front++
		rear++
	}
	return ans
}
