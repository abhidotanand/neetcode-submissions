import "slices"
func topKFrequent(nums []int, k int) []int {
	var n int = len(nums)
	var m map[int]int = make(map[int]int)
	var tmp [][]int = make([][]int, n+1)
	var ans []int = make([]int, 0)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}

	for _, num := range nums {
		if _, ok := m[num]; ok {
			if !slices.Contains(tmp[m[num]], num) {
				tmp[m[num]] = append(tmp[m[num]], num)
			}
		}
	}
	cnt := 0
	for i := n; i > 0; i-- {
		if len(tmp[i]) != 0 {
			for _, num := range tmp[i]{
				ans = append(ans, num)
				cnt++
				if cnt == k{
					return ans
				}
			}
		}
	}
	return ans
}
