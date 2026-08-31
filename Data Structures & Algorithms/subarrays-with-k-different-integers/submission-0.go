func subarraysWithKDistinct(nums []int, k int) int {
    n := len(nums)
    count := make([]int, n+1)
    res, l, cnt := 0, 0, 0

    for r := 0; r < n; r++ {
        count[nums[r]]++
        if count[nums[r]] == 1 {
            k--
        }

        if k < 0 {
            count[nums[l]]--
            l++
            k++
            cnt = 0
        }

        if k == 0 {
            for count[nums[l]] > 1 {
                count[nums[l]]--
                l++
                cnt++
            }
            res += cnt + 1
        }
    }

    return res
}