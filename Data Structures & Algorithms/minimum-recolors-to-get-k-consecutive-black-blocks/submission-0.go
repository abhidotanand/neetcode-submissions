func minimumRecolors(blocks string, k int) int {
	var n int = len(blocks)
	var ans, tmp int

	for i := 0; i < k; i++ {
		if string(blocks[i]) == "W" {
			tmp++
		}
	}

	ans = tmp
	front := 0
	rear := k

	for rear < n {
		if string(blocks[rear]) == "W" {
			tmp++
		}

		if string(blocks[front]) == "W" {
			tmp--
		}
		ans = min(ans, tmp)
		front++
		rear++
	}
	return ans
}
