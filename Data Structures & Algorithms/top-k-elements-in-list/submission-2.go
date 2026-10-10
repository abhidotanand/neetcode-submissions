func topKFrequent(nums []int, k int) []int {
	var n int = len(nums)
	var m map[int]int = make(map[int]int)
	var freq_arr [][]int = make([][]int, n+1)
	var ans []int = make([]int, 0)

	for i := 0; i < n; i++ {
		m[nums[i]]++
		freq_arr[m[nums[i]]] = append(freq_arr[m[nums[i]]], nums[i])
		if m[nums[i]] > 1 {
			freq_arr[m[nums[i]] - 1] = freq_arr[m[nums[i]] - 1][:len(freq_arr[m[nums[i]] - 1])-1]
		}
	}
	
	counter := 0
	for i := n; i > 0; i-- {
		if len(freq_arr[i]) == 0 {
			continue
		} else {
			for counter < k && len(freq_arr[i]) > 0 {
				ans = append(ans, freq_arr[i][0])
				freq_arr[i] = freq_arr[i][1:]
				counter++
			}
		}
		if counter == k {
			break
		}
	}
	return ans
}
