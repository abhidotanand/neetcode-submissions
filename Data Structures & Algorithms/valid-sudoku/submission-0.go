import "maps"
func isValidSudoku(board [][]byte) bool {
	var n int = 9
	var nums_map map[byte]struct{} = map[byte]struct{}{'1':struct{}{}, '2':struct{}{}, '3':struct{}{}, '4':struct{}{}, '5':struct{}{}, '6':struct{}{}, '7':struct{}{}, '8':struct{}{}, '9':struct{}{}}

	for i := 0; i < n; i++ {
		tmp_map := maps.Clone(nums_map)
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			} else if _,ok := tmp_map[board[i][j]]; ok {
				delete(tmp_map, board[i][j])
			} else {
				return false
			}
		}
	}

	for j := 0; j < n; j++ {
		tmp_map := maps.Clone(nums_map)
		for i := 0; i < n; i++ {
			if board[i][j] == '.' {
				continue
			} else if _,ok := tmp_map[board[i][j]]; ok {
				delete(tmp_map, board[i][j])
			} else {
				return false
			}
		}
	}

	for r := 0; r < n; r = r + 3 {
		for c := 0; c < n; c = c + 3 {
			tmp_map := maps.Clone(nums_map)
			for i := r; i < r +3; i++ {
				for j := c; j < c + 3; j++ {
					if board[i][j] == '.' {
						continue
					} else if _,ok := tmp_map[board[i][j]]; ok {
						delete(tmp_map, board[i][j])
					} else {
						return false
					}
				}
			}
		}
	}
	return true
}
