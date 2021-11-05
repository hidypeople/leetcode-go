package task0441

import "math"

// You have n coins and you want to build a staircase with these coins.
// The staircase consists of k rows where the ith row has exactly i coins.
// The last row of the staircase may be incomplete.
// Given the integer n, return the number of complete rows of the staircase you will build.
//
// Constraints:
//   1 <= n <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Arranging Coins.
//   Memory Usage: 2.2 MB, less than 43.28% of Go online submissions for Arranging Coins.
func arrangeCoins(n int) int {
	return int(math.Floor(-0.5 + math.Sqrt(float64(1+8*n))/2))
}
