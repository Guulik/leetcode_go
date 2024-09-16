package SlidingWindow

func BestTimeToTrade(prices []int) int {
	minimal := 10000
	maxProfit := 0
	for _, p := range prices {
		if p < minimal {
			minimal = p
		}
		if p-minimal > maxProfit {
			maxProfit = p - minimal
		}
	}
	return maxProfit
}
