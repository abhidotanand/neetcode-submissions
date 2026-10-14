func majorityElement(nums []int) int {
    var n int = len(nums)
	var candi, cnt int

	for i := 0; i < n; i++ {
		if cnt == 0 {
			candi = nums[i]
			cnt++
		} else if candi == nums[i] {
			cnt++
		} else if candi != nums[i] {
			cnt--
		}
	}
	return candi
}
