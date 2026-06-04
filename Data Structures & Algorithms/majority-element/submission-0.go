func majorityElement(nums []int) int {
	var n int = len(nums)
    var cand int = -1e9-1
	var cnt int

	for i := 0; i < n; i++ {
		if cnt == 0 {
			cand = nums[i]
			cnt = 1
			continue
		}
		if nums[i] != nums[i-1] {
			cnt--
		} else {
			cnt++
		}
	}
	return cand
}