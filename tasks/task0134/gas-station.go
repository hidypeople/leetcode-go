package task0134

// There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith
// station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.
// Given two integer arrays gas and cost, return the starting gas station's index if you can travel
// around the circuit once in the clockwise direction, otherwise return -1. If there exists a
// solution, it is guaranteed to be unique
//
// Constraints:
//   gas.length == n
//   cost.length == n
//   1 <= n <= 10^5
//   0 <= gas[i], cost[i] <= 10^4
//
// Results:
//   Runtime: 48 ms, faster than 78.57% of Go online submissions for Gas Station.
//   Memory Usage: 9.6 MB, less than 37.76% of Go online submissions for Gas Station.
func canCompleteCircuit(gasArray []int, costArray []int) int {
	n, grandTotal, currentTotal, startIdx := len(gasArray), 0, 0, 0
	for i := 0; i < n; i++ {
		delta := gasArray[i] - costArray[i]
		currentTotal += delta
		grandTotal += delta
		if currentTotal < 0 {
			startIdx = i + 1
			currentTotal = 0
		}
	}
	if grandTotal < 0 {
		return -1
	} else {
		return startIdx
	}
}
