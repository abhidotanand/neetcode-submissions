func simplifyPath(path string) string {
	var deq []string
	var dirs []string = strings.Split(path, "/")
	var n int = len(dirs)
	var ans string

	for i := 0; i < n; i++ {
		if strings.Contains(dirs[i], ".") {
			if strings.Count(dirs[i], ".") > 2 && len(dirs[i]) == strings.Count(dirs[i], ".") {
				deq = append(deq, dirs[i])
			} else if strings.Count(dirs[i], ".") == 2 && len(dirs[i]) == 2 {
				if len(deq) > 0 {
					deq = deq[:len(deq)-1]
				}
			} else if strings.Count(dirs[i], ".") == 1 && len(dirs[i]) == 1 {
				continue
			} else {
				if len(strings.TrimSpace(dirs[i])) > 0 {
					deq = append(deq, dirs[i])
				}
			}
		} else {
			if len(strings.TrimSpace(dirs[i])) > 0 {
				deq = append(deq, dirs[i])
			}
		}
		fmt.Println(deq)
	}
	
	for _, w := range(deq){
		ans += "/" + w
	}
	
	if len(ans) == 0 {
		return "/"
	}

	return ans
}
