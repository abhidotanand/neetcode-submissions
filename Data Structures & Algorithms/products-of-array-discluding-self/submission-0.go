func productExceptSelf(nums []int) []int {
	var n int = len(nums)
	var left_prod []int = make([]int, n)
	var right_prod []int = make([]int, n)
	var ans []int = make([]int, n)

	left_prod[0] = nums[0]
	right_prod[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		left_prod[i] = left_prod[i-1]*nums[i]
	}

	for i := n-2; i >= 0; i-- {
		right_prod[i] = right_prod[i+1]*nums[i]
	}
	
	fmt.Println(left_prod, right_prod)

	ans[0] = right_prod[1]
	ans[n-1] = left_prod[n-2]

	for i := 1; i < n-1; i++ {
		ans[i] = left_prod[i-1]*right_prod[i+1]
	}
	return ans
}
