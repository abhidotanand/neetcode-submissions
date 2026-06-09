func maxProfit(prices []int) int {
	var n int = len(prices)
	var bought_at int = prices[0]
	var profit int
 	for i := 1; i < n; i++ {
		if prices[i] > bought_at {
			profit += prices[i] - bought_at
		}
		bought_at = prices[i]
	}
	return profit
}