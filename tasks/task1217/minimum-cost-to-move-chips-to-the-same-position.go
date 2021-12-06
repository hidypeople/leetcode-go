package task1217

// We have n chips, where the position of the ith chip is position[i].
// We need to move all the chips to the same position. In one step, we can change the
// position of the ith chip from position[i] to:
//   position[i] + 2 or position[i] - 2 with cost = 0.
//   position[i] + 1 or position[i] - 1 with cost = 1.
// Return the minimum cost needed to move all the chips to the same position.
//
// Constraints:
//   1 <= position.length <= 100
//   1 <= position[i] <= 10^9
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Cost to Move Chips to The Same Position.
//   Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Minimum Cost to Move Chips to The Same Position.
func minCostToMoveChips(position []int) int {
	odds, evens := 0, 0
	for _, pos := range position {
		if pos%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	if evens > odds {
		return odds
	} else {
		return evens
	}
}
