func subarraySum(nums []int, k int) int {
	var n int = len(nums)
	var hash map[int]int = make(map[int]int)
	var freq []int = make([]int, n)
	var ans int

	freq[0] = nums[0]
	hash[0] = 1

	for i := 1; i < n; i++ {
		freq[i] = freq[i-1] + nums[i]
	}

	for i := 0; i < n; i++ {
		if _, ok := hash[freq[i] - k]; ok {
			ans += hash[freq[i] - k]
		}
		if _, ok := hash[freq[i]]; ok {
			hash[freq[i]]++
		} else {
			hash[freq[i]] = 1
		}
	}

	return ans
}
