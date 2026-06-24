func decodeString(s string) string {
	var st_num []int
	var st_chr []string
	var ans string

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if num, err := strconv.Atoi(c); err == nil {
			st_num = append(st_num, num)
		} else if c == "[" {
			st_chr = append(st_chr, "")
		} else if c == "]" {
			tmp := strings.Repeat(st_chr[len(st_chr)-1], st_num[len(st_num)-1])
			st_chr = st_chr[:len(st_chr)-1]
			st_num = st_num[:len(st_num)-1]
			if len(st_chr) == 0 {
				ans += tmp
			} else {
				st_chr[len(st_chr)-1] += tmp
			}
		} else {
			if len(st_num) == 0 {
				st_num = append(st_num, 1)
				st_chr = append(st_chr, c)
			} else {
				st_chr[len(st_chr)-1] += c
			}
		}
	}
	return ans + st_chr[0]
}
