package task0739

// Given an array of integers temperatures represents the daily temperatures,
// return an array answer such that answer[i] is the number of days you have
// to wait after the ith day to get a warmer temperature. If there is no future
// day for which this is possible, keep answer[i] == 0 instead.
//
// Constraints:
//   1 <= temperatures.length <= 10^5
//   30 <= temperatures[i] <= 100
//
// Results:
//   Runtime: 132 ms, faster than 98.17% of Go online submissions for Daily Temperatures.
//   Memory Usage: 9.2 MB, less than 90.85% of Go online submissions for Daily Temperatures.
func dailyTemperatures(temperature []int) []int {
	result := make([]int, len(temperature))
	stack := []int{}
	for i, t := range temperature {
		for len(stack) > 0 && temperature[stack[len(stack)-1]] < t {
			// For all previous temperatures that lesser than current temperature:
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[j] = i - j
		}

		// Stack contains previous temperatures:
		stack = append(stack, i)
	}
	return result
}
