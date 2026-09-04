func findMissingAndRepeatedValues(grid [][]int) []int {
    N := len(grid)
    seen := make(map[int]bool)
    doubleVal, missing := 0, 0

    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            if seen[grid[i][j]] {
                doubleVal = grid[i][j]
            }
            seen[grid[i][j]] = true
        }
    }

    for num := 1; num <= N*N; num++ {
        if !seen[num] {
            missing = num
            break
        }
    }

    return []int{doubleVal, missing}
}