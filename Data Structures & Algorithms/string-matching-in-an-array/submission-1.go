import (
	"slices"
	"cmp"
)
func stringMatching(words []string) []string {
	var ans []string
	var n int = len(words)
	var found bool

    slices.SortFunc(words, func (a, b string) int {
		return cmp.Compare(len(a), len(b))
	})

	fmt.Println(words)
	for i := 0; i < n-1; i++ {
		found = false
		for j := i+1; j < n; j++ {
			front := 0
			rear := len(words[i])
			for rear <= len(words[j]) {
				if words[i] == words[j][front:rear] {
					ans = append(ans, words[i])
					found = true
					break
				}
				front++
				rear++
			}
			if found {
				break
			}
		}
	}
	return ans
}