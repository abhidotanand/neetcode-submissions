func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var n int = len(nums1)
    var m int = len(nums2)
    var ms, ans []int
    var ht map[int]int = make(map[int]int)

    for i := 0; i < m; i++ {
        if len(ms) == 0 || nums2[i] <= ms[len(ms)-1] {
            ms = append(ms, nums2[i])
        } else {
            for len(ms) > 0 && ms[len(ms)-1] < nums2[i] {
                ht[ms[len(ms)-1]] = nums2[i]
                ms = ms[:len(ms)-1]
            }
            ms = append(ms, nums2[i])
        }
    }

    for i := 0; i < n; i++ {
        if val, ok := ht[nums1[i]]; ok {
            ans = append(ans, val)
        } else {
            ans = append(ans, -1)
        }
    }
    
    return ans
}
