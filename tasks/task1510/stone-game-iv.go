package task1510

import "math"

// Alice and Bob take turns playing a game, with Alice starting first.
// Initially, there are n stones in a pile. On each player's turn, that
// player makes a move consisting of removing any non-zero square number of stones in the pile.
// Also, if a player cannot make a move, he/she loses the game.
// Given a positive integer n, return true if and only if Alice wins the
// game otherwise return false, assuming both players play optimally.
//
// Constraints:
//   1 <= n <= 10^5
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for Stone Game IV.
//   Memory Usage: 3.3 MB, less than 77.78% of Go online submissions for Stone Game IV.
func winnerSquareGame(n int) bool {
	knownResults := make([]int, n+1)
	return winner(n, &knownResults)
}

func winner(n int, knownResults *[]int) bool {
	knownResult := (*knownResults)[n]
	if knownResult != 0 {
		if knownResult > 0 {
			return true
		} else {
			return false
		}
	}
	for i := int(math.Sqrt(float64(n))); i > 0; i-- {
		stonesLeft := n - i*i
		if stonesLeft == 0 || !winner(stonesLeft, knownResults) {
			(*knownResults)[n] = 1
			return true
		}
	}
	(*knownResults)[n] = -1
	return false
}
