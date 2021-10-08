package task0121

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a
// different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
// Constraints:
//   1 <= prices.length <= 10^5
//   0 <= prices[i] <= 10^4
//
// Results:
//   Runtime: 116 ms, faster than 95.68% of Go online submissions for Best Time to Buy and Sell Stock.
//   Memory Usage: 8.3 MB, less than 96.59% of Go online submissions for Best Time to Buy and Sell Stock.
func maxProfit(prices []int) int {
	result := 0
	currentMin := prices[0]
	var price int
	for i := 1; i < len(prices); i++ {
		price = prices[i]
		if currentMin > price {
			currentMin = price
		} else if currentMin < price && result < price-currentMin {
			result = price - currentMin
		}
	}
	return result
}
