func maxProfit(prices []int) int {
	var n int = len(prices)
	var ans, tmp int = 0, prices[0]

	for i := 1; i < n; i++ {
		if prices[i] <= tmp {
			tmp = prices[i]
		} else {
			ans = max(ans, max(0, prices[i]-tmp))
		}
	}
	return ans
}
