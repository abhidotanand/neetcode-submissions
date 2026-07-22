import (
	"slices"
	"cmp"
)
func stringMatching(words []string) []string {
	var ans []string
	var n int = len(words)
    slices.SortFunc(words, func (a, b string) int {
		return cmp.Compare(len(a), len(b))
	})
	fmt.Println(words)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			front := 0
			rear := len(words[i])
			for rear <= len(words[j]) {
				if words[i] == words[j][front:rear] {
					ans = append(ans, words[i])
				}
				front++
				rear++
			}
		}
	}
	return ans
}