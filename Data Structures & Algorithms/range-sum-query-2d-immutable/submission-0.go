type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var n int = len(matrix)
	var m int = len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	for j := 0; j < m; j++ {
		for i := 1; i < n; i++ {
			matrix[i][j] += matrix[i-1][j]
		}
	}
	return NumMatrix{
		matrix: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.matrix[row2][col2]
	} else if row1 == 0 {
		return this.matrix[row2][col2] - this.matrix[row2][col1-1]
	} else if col1 == 0 {
		return this.matrix[row2][col2] - this.matrix[row1-1][col2]
	}
	return this.matrix[row2][col2] - this.matrix[row1-1][col2] - this.matrix[row2][col1-1] + this.matrix[row1-1][col1-1]
}


// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
