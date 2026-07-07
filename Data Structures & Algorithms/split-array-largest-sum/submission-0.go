import "slices"
func splittable(nums []int, summ int, k int) bool {
	var n int = len(nums)
	var cnt int = 1
	var ptr int

	tmp := 0
	for ptr < n {
		tmp += nums[ptr]
		if tmp > summ {
			cnt++
			tmp = 0
		} else {
			ptr++
		}
	}

	return cnt <= k
}
func splitArray(nums []int, k int) int {
	var left, right int = slices.Max(nums), 0

	for i := range nums {
		right += nums[i]
	}

	var ans int = right

	for left <= right {
		mid := (left + right)/2

		if splittable(nums, mid, k) {
			ans = min(ans, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
