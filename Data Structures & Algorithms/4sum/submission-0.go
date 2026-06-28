import "slices"
func fourSum(nums []int, target int) [][]int {
	var n int = len(nums)
	var i, j, k, l int = 0, 1, 2, n-1
	var ans [][]int = make([][]int, 0)

	slices.Sort(nums)

	for i < n {
		for i > 0 && i < n && nums[i] == nums[i-1] {
			i++
		}
		
		j = i+1
		for j < k {
			for j < k && nums[j] == nums[j-1] {
				j++
			}

			k = j+1
			for k < l {
				if nums[i] + nums[j] + nums[k] + nums[l] == target{
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if nums[i] + nums[j] + nums[k] + nums[l] < target {
					k++
				} else {
					l--
				}
			}
			j++
		}
		i++
	}
	return ans
}
