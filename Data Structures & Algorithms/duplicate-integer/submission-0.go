func hasDuplicate(nums []int) bool {
    var int_set map[int]struct{} = make(map[int]struct{})

	for _, num := range nums{
		if _, ok := int_set[num]; ok{
			return true
		} else {
			int_set[num] = struct{}{}
		}
	}
	return false
}
