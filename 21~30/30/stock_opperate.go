func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max, lowest := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[lowest] {
			lowest = i
		} else if cur := prices[i] - prices[lowest]; cur > max {
			max = cur
		}
	}
	return max
}