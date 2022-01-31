package task1672

// You are given an m x n integer grid accounts where accounts[i][j] is the amount of money
// the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.
// A customer's wealth is the amount of money they have in all their bank accounts.
// The richest customer is the customer that has the maximum wealth.
//
// Constraints:
//   m == accounts.length
//   n == accounts[i].length
//   1 <= m, n <= 50
//   1 <= accounts[i][j] <= 100
//
// Results:
//   Runtime: 2 ms, faster than 94.53% of Go online submissions for Richest Customer Wealth.
//   Memory Usage: 3.1 MB, less than 99.50% of Go online submissions for Richest Customer Wealth.
func maximumWealth(accounts [][]int) int {
	max := 0
	var curr int
	for _, account := range accounts {
		curr = 0
		for _, x := range account {
			curr += x
		}
		if curr > max {
			max = curr
		}
	}
	return max
}
