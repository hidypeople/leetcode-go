package task0123

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// Find the maximum profit you can achieve. You may complete at most two transactions.
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
// Constraints:
//   1 <= prices.length <= 10^5
//   0 <= prices[i] <= 10^5
//
// Results:
//   Runtime: 104 ms, faster than 99.61% of Go online submissions for Best Time to Buy and Sell Stock III.
//   Memory Usage: 9.1 MB, less than 84.71% of Go online submissions for Best Time to Buy and Sell Stock III.
func maxProfit(prices []int) int {
	// Common DP formula for k-transactions on i-th day:
	// dp[k, i] = max(dp[k, i-1], prices[i] - prices[j] + dp[k-1, j-1]), j=[0..i-1]
	//   1. dp[k-1, j-1] - previous transaction max profit
	//   2. prices[i] - prices[j] - current transaction max profit
	//   3. dp[k, i-1] - max profit for previous day

	// The idea is simple: buy and sell variables represents how much money do we have
	// in our wallet after 1st or 2nd buy and sell operation. All operations are going
	// one-by-one buy1 -> sell1 -> buy2 -> sell2
	buy1, buy2, sell1, sell2 := -1<<63, -1<<63, 0, 0
	for _, price := range prices {
		buy1 = max(buy1, -price)       // Spend money
		sell1 = max(sell1, buy1+price) // Take profit - This is the maximum profit of one single deal
		buy2 = max(buy2, sell1-price)  // Spend money second time
		sell2 = max(sell2, buy2+price) // Take profit second time
	}
	return sell2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
