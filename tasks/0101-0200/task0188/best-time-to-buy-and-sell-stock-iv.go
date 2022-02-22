package task0188

// You are given an integer array prices where prices[i] is the price of a given stock on the ith day,
// and an integer k.
// Find the maximum profit you can achieve. You may complete at most k transactions.
// Note: You may not engage in multiple transactions simultaneously
// (i.e., you must sell the stock before you buy again).
//
// Constraints:
//   0 <= k <= 100
//   0 <= prices.length <= 1000
//   0 <= prices[i] <= 1000
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock IV.
//   Memory Usage: 2 MB, less than 98.75% of Go online submissions for Best Time to Buy and Sell Stock IV.
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) < 1 {
		return 0
	}
	MIN := -1 << 63
	buy, sell := make([]int, k), make([]int, k)
	for i := 0; i < k; i++ {
		buy[i] = MIN
	}

	for _, price := range prices {
		buy[0] = max(buy[0], -price)
		sell[0] = max(sell[0], buy[0]+price)
		for i := 1; i < k; i++ {
			buy[i] = max(buy[i], sell[i-1]-price)
			sell[i] = max(sell[i], buy[i]+price)
		}
	}

	return sell[k-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
