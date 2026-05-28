func twoSum(nums []int, target int) []int {
    var ans []int
    var m = make(map[int]int)
    for ind, num := range nums{
        m[num] = ind
    }

    for ind, num := range nums{
        if _, ok := m[target - num]; ok && ind != m[target - num] {
            ans = append(ans, ind)
            ans = append(ans, m[target - num])
            break
        }
    }
    return ans
}