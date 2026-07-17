func replaceElements(arr []int) []int {
	var n int = len(arr)

	tmp := 0
	for i := n-2; i >= 0; i-- {
		tmp1 := arr[i]
		arr[i] = max(arr[i+1], tmp)
		tmp = tmp1
	}
	
	arr[n-1] = -1

	return arr
}
