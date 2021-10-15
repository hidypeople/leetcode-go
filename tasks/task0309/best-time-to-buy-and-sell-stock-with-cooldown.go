package task0309

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// Find the maximum profit you can achieve. You may complete as many transactions as you like
// (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
// After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the
// stock before you buy again).
//
// Constraints:
//   1 <= prices.length <= 5000
//   0 <= prices[i] <= 1000
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock with Cooldown.
//   Memory Usage: 2.4 MB, less than 59.18% of Go online submissions for Best Time to Buy and Sell Stock with Cooldown.
func maxProfit(prices []int) int {
	sold, hold, rest := 0, -1<<63, 0
	for i := 0; i < len(prices); i++ {
		prvSold := sold
		sold = hold + prices[i]
		hold = max(hold, rest-prices[i])
		rest = max(rest, prvSold)
	}
	return max(sold, rest)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
