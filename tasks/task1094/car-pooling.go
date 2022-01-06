package task1094

// There is a car with capacity empty seats. The vehicle only drives east (i.e., it cannot turn around and drive west).
// You are given the integer capacity and an array trips where trip[i] = [numPassengersi, fromi, toi] indicates that
// the ith trip has numPassengersi passengers and the locations to pick them up and drop them off are fromi and toi
// respectively. The locations are given as the number of kilometers due east from the car's initial location.
// Return true if it is possible to pick up and drop off all passengers for all the given trips, or false otherwise.
//
// Constraints:
//   1 <= trips.length <= 1000
//   trips[i].length == 3
//   1 <= numPassengersi <= 100
//   0 <= fromi < toi <= 1000
//   1 <= capacity <= 10^5
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for Car Pooling.
//   Memory Usage: 3.5 MB, less than 70.83% of Go online submissions for Car Pooling.
func carPooling(trips [][]int, capacity int) bool {
	currentPassengers := make([]int, 1001)
	var tripLen int
	// store for each kilometer "the amount of changed passengers"
	for _, trip := range trips {
		tripLen = trip[0]
		currentPassengers[trip[1]] += tripLen
		currentPassengers[trip[2]] -= tripLen
	}
	curr := 0
	for _, currentPassenger := range currentPassengers {
		curr += currentPassenger
		if curr > capacity {
			return false
		}
	}
	return true
}
