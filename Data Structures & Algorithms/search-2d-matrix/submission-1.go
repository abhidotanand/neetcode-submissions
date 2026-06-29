func searchMatrix(matrix [][]int, target int) bool {
	var n, m int = len(matrix), len(matrix[0])
	var top, bottom, left, right, row, col int = 0, n-1, 0, m-1, 0, 0

	for top <= bottom {
		row = (top + bottom)/2
		if matrix[row][m-1] == target {
			return true
		} else if matrix[row][m-1] < target {
			top = row + 1
		} else {
			bottom = row - 1
		}
	}
	if matrix[row][m-1] < target && row < n-1 {
		row += 1
	}

	for left <= right {
		col = (left + right)/2
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = col + 1
		} else {
			right = col - 1
		}
	}
	return false
}
