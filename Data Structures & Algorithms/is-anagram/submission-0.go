func isAnagram(s string, t string) bool {
	var chr_arr [26]int

	for _, c := range s{
		chr_arr[rune(c)-97]++
	}

	for _, c := range t{
		chr_arr[rune(c)-97]--
	}

	for _, c := range chr_arr{
		if c != 0{
			return false
		}
	}
	return true
}
