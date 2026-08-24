func getRow(rowIndex int) []int {
    row := []int{1}
    for i := 1; i <= rowIndex; i++ {
        row = append(row, row[len(row)-1]*(rowIndex-i+1)/i)
    }
    return row
}