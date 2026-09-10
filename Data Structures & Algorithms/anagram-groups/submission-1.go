func groupAnagrams(strs []string) [][]string {
	var hash_arr [][26]int = make([][26]int, 0)
	var tmp_hash map[[26]int][]int = make(map[[26]int][]int)
	var ans [][]string = make([][]string, 0)
	var wc [26]int

	for _, str := range strs {
		for _, l := range str {
			wc[rune(l)-97]++
		}
		hash_arr = append(hash_arr, wc)
		wc = [26]int{}
	}

	for i, ele := range hash_arr {
		if _, ok := tmp_hash[ele]; ok {
			tmp_hash[ele] = append(tmp_hash[ele], i)
		} else {
			tmp_hash[ele] = []int{i}
		}
	}

	for i := range hash_arr{
		if _,ok := tmp_hash[hash_arr[i]]; ok {
			tmp := []string{}
			for _, ele := range tmp_hash[hash_arr[i]] {
				tmp = append(tmp, strs[ele])
			}
			delete(tmp_hash, hash_arr[i])
			ans = append(ans, tmp)
			tmp = nil
		}
	}
	return ans
}
