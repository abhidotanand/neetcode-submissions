func hasDuplicate(nums []int) bool {
    var n int = len(nums)
    var m map[int]struct{} = make(map[int]struct{})

    for i := 0; i < n; i++ {
        if _, ok := m[nums[i]]; ok {
            return true
        } else {
            m[nums[i]] = struct{}{}
        }
    }
    return false
}
