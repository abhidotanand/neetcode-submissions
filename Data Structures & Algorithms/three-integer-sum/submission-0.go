import "slices"
func threeSum(nums []int) [][]int {
	var n int = len(nums)
	var i,j,k int
	var ans [][]int = make([][]int, 0)
	
	slices.Sort(nums)

	for i < n-1 {
		for i > 0 && i < n && nums[i] == nums[i-1] {
			i++
		}

		j = i+1
		k = n-1

		for j < k {
			if nums[i] + nums[j] + nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < n && nums[j] == nums[j-1] {
					j++
				}
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[i] + nums[j] + nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
		i++
	}
	return ans
}
