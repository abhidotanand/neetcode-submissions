func numUniqueEmails(emails []string) int {
    var m map[string]struct{} = make(map[string]struct{})

	for _, email := range emails {
		tmp := ""
		at := false
		plus := false

		for e := range email {
			if string(email[e]) == "+" {
				plus = true
				continue
			}
			if string(email[e]) == "@" {
				at = true
				continue
			}
			
			if at {
				tmp += string(email[e])
			} else {
				if plus {
					continue
				} else {
					if string(email[e]) == "." {
						continue
					} else {
						tmp += string(email[e])
					}
				}
			}
		}
		m[tmp] = struct{}{}
	}
	fmt.Println(m)
	return len(m)
}