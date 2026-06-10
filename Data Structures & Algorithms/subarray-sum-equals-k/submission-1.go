func subarraySum(nums []int, k int) int {
	var n int = len(nums)
	var freq_sum []int = make([]int, n)
	var freq_hash map[int]int = make(map[int]int)
	var ans int

	freq_sum[0] = nums[0]
	freq_hash[0] = 1

	for i := 1; i < n; i++ {
		freq_sum[i] = freq_sum[i-1] + nums[i]
	}
	fmt.Println(freq_sum)
	for i := 0; i < n; i++ {
		if _, ok := freq_hash[freq_sum[i]]; !ok {
			freq_hash[freq_sum[i]] = 1
		} else {
			freq_hash[freq_sum[i]]++
		}
	}

	for i := 0; i < n; i++ {
		if _,ok := freq_hash[freq_sum[i] - k]; ok{
			ans += freq_hash[freq_sum[i] - k]
		}
	}

	return ans
}
