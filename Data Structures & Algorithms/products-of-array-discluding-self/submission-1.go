// import "slices"
func productExceptSelf(nums []int) []int {
	var n int = len(nums)
	var left_prod, right_prod, ans []int = make([]int, n), make([]int, n), make([]int, n)

	left_prod[0] = 1
	right_prod[n-1] = 1

	for i := 1; i < n; i++ {
		left_prod[i] = left_prod[i-1]*nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		right_prod[i] = right_prod[i+1]*nums[i+1]
	}
	
	ans[0] = right_prod[0]
	ans[n-1] = left_prod[n-1]
	
	for i := 1; i < n-1; i++ {
		ans[i] = left_prod[i]*right_prod[i]
	}

	return ans
}
