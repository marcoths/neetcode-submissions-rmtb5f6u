func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxTrade := 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		minPrice = min(price, minPrice)
		trade := price - minPrice
		maxTrade = max(trade, maxTrade)
	}

	return maxTrade
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

