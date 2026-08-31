func sortedSquares(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    l, r := 0, n-1
    resIndex := n - 1

    for l <= r {
        if abs(nums[l]) > abs(nums[r]) {
            res[resIndex] = nums[l] * nums[l]
            l++
        } else {
            res[resIndex] = nums[r] * nums[r]
            r--
        }
        resIndex--
    }

    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}