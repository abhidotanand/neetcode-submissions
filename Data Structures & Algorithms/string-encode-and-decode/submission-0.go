type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// strs = nil
	// strs = []string{"Abhi", "  bh  j  " , "  ", ""}
	// strs = []string{"" , "", ""}
	if len(strs) == 0 {
		return ""
	}
	var encoded_string string = ""

	for _, str := range strs {
		if len(str) == 0{
			encoded_string += "0\t\n"
			continue
		}
		for _, s := range str {
			encoded_string += fmt.Sprintf("%d\t",s)
		}
		n := len(encoded_string)
		if n == 0{
			encoded_string += "\n"
		} else{ 
			encoded_string = encoded_string[0:n-1] + "\n"
		}
	}
	return encoded_string[0:len(encoded_string)-1]
}


func (s *Solution) Decode(encoded string) []string {
	var ans []string =  make([]string, 0)
	if len(encoded) == 0{
		return ans
	}
	encoded_words := strings.Split(encoded, "\n")
	for _, ew := range encoded_words {
		tmp := strings.Split(ew, "\t")
		tmp_str := ""
		for _, l := range tmp {
			lt, _ := strconv.Atoi(l)
			if lt == 0 {
				tmp_str = ""
			} else {
				tmp_str += fmt.Sprintf("%c", rune(lt))
			}
		}
		fmt.Println(tmp_str)
		ans = append(ans, tmp_str)
	}
	return ans
}
