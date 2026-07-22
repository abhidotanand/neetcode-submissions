func generate(numRows int) [][]int {
	var ans [][]int = make([][]int, numRows)

	for row := range numRows {
		ans[row] = make([]int, row + 1)
		ans[row][0] = 1
		ans[row][len(ans[row]) - 1] = 1
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < len(ans[i]) - 1; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}
