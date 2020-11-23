func maxProfit(prices []int) int {
	profits := 0
	stack := prices[0]
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] > prices[i+1] {
			profits += prices[i] - stack
			stack = prices[i+1]
		}
	}
	profits += prices[len(prices)-1] - stack
	return profits
}