func majorityElement(nums []int) []int {
	var n int = len(nums)
	var candi_one, candi_two, candi_one_sum, candi_two_sum, cnt_candi_one, cnt_candi_two int

	for i := 0; i < n; i++ {
		if candi_one_sum <= 0 && nums[i] != candi_two {
			candi_one = nums[i]
			candi_one_sum = 1
		} else if candi_two_sum <= 0 && nums[i] != candi_one{
			candi_two = nums[i]
			candi_two_sum = 1
		} else if nums[i] == candi_one {
			candi_one_sum++
		}  else if nums[i] == candi_two {
			candi_two_sum++
		} else {
			candi_one_sum--
			candi_two_sum--
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] == candi_one && candi_one_sum > 0 {
			cnt_candi_one++
		} else if nums[i] == candi_two && candi_two_sum > 0 {
			cnt_candi_two++
		}
	}

	if cnt_candi_one > n/3 && cnt_candi_two > n/3 {
		return []int{candi_one, candi_two}
	} else if cnt_candi_one > n/3 {
		return []int{candi_one}
	} else if cnt_candi_two > n/3 {
		return []int{candi_two}
	}
	return []int{}
}
