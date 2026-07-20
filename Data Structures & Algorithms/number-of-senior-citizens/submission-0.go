func countSeniors(details []string) int {
	var n int = len(details)
	var ans int

	for i := 0; i < n; i++ {
		if age, _ := strconv.Atoi(string(details[i][11]) + string(details[i][12])); age > 60 {
			ans++
		}
	}
	return ans
}