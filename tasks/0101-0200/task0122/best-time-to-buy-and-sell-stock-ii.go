package task0122

// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the
// stock at any time. However, you can buy it then immediately sell it on the same day.
// Find and return the maximum profit you can achieve.
//
// Constraints:
//   1 <= prices.length <= 3 * 10^4
//   0 <= prices[i] <= 10^4
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock II.
//   Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Best Time to Buy and Sell Stock II.
func maxProfit(prices []int) int {
	result := 0
	var diff int
	for i, val := range prices {
		if i > 0 {
			diff = val - prices[i-1]
			if diff >= 0 {
				result += diff
			}
		}
	}
	return result
}
